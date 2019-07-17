package command

import (
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/NoahOrberg/nimvle.nvim/nimvle"
	"github.com/neovim/go-client/nvim"
)

var (
	initializedNowPlaying = make(chan struct{}, 1)
)

type Spotify struct {
	NowPlaying string
	n          *nimvle.Nimvle
	m          *sync.Mutex

	// For initialize GetNowPlaying
	onceRecv *sync.Once
	onceSnd  *sync.Once
}

func (s *Spotify) init() {
	// initialize
	s.m = new(sync.Mutex)
	s.onceRecv = new(sync.Once)
	s.onceSnd = new(sync.Once)
}

func NewSpotify() *Spotify {
	s := &Spotify{}
	s.init()
	return s
}

func (s *Spotify) Init(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	s.n = nimvle
	s.pollingNowPlaying()

	nowPlaying, err := s.getNowPlaying(
		s.runtimePath(nimvle))
	if err != nil {
		nimvle.Log(err)
	}

	s.m.Lock()
	defer s.m.Unlock()
	s.NowPlaying = nowPlaying

	// NOTE: At first, GetNowPlaying method wait this channel
	s.onceSnd.Do(func() {
		initializedNowPlaying <- struct{}{}
	})

	return nil
}

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	s.onceRecv.Do(func() {
		<-initializedNowPlaying
	})
	return s.NowPlaying, nil
}

func (s *Spotify) getNowPlaying(rtp string) (string, error) {
	cmd := exec.Command(
		"/usr/bin/osascript", rtp+"spotify_util/now_playing.applescript")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out[:len(out)-1]), nil // drop ^@
}

func (s *Spotify) pollingNowPlaying() {
	rtp := s.runtimePath(s.n)
	go func() {
		for {
			nowPlaying, err := s.getNowPlaying(rtp)
			if err != nil {
				continue
			}

			if nowPlaying != s.NowPlaying {
				s.m.Lock()
				s.NowPlaying = nowPlaying
				s.m.Unlock()
			}

			time.Sleep(1500 * time.Millisecond)
		}
	}()
}

func (s *Spotify) runtimePath(nimvle *nimvle.Nimvle) string {
	irtp, err := nimvle.Eval(`g:ayunis_rtp`)
	if err != nil {
		nimvle.Log(err)
		panic(err)
	}

	srtp := irtp.(string)
	if !strings.HasSuffix(srtp, "/") {
		srtp = strings.Join([]string{srtp, "/"}, "")
	}

	return srtp
}

func (s *Spotify) exec(v *nvim.Nvim, cmd string, args ...string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	// StupidSolution: Spotify.app takes time to be reflected. So polling it
	limit := 100
	for i := 0; i < limit; i++ {
		// NOTE: check about it is changing song
		if strings.Contains(args[0], "next") ||
			strings.Contains(args[0], "shuffle") {

			nowPlaying, err := s.getNowPlaying(
				s.runtimePath(nimvle))
			if err != nil {
				nimvle.Log(err.Error())
			}
			if nowPlaying != s.NowPlaying {
				s.m.Lock()
				defer s.m.Unlock()

				s.NowPlaying = nowPlaying
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}

func (s *Spotify) Next(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/playback_next.applescript")
}

func (s *Spotify) Prev(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/playback_prev.applescript")
}

func (s *Spotify) Toggle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/playback_toggle.applescript")
}

func (s *Spotify) ToggleRepeat(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/toggle_repeat.applescript")
}

func (s *Spotify) ToggleShuffle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/toggle_shuffle.applescript")
}

func (s *Spotify) VolumeUp(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/volume_up.applescript")
}

func (s *Spotify) VolumeDown(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, "/usr/bin/osascript",
		s.runtimePath(nimvle)+"spotify_util/volume_down.applescript")
}
