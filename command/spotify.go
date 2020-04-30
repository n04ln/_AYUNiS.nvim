package command

import (
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/n04ln/nimvle.nvim/nimvle"

	"github.com/neovim/go-client/nvim"
)

var (
	initializedNowPlaying = make(chan struct{}, 1)
)

const (
	getNowPlayingPollingInterval = 3000 * time.Millisecond
	appleScriptCmd               = "/usr/bin/osascript"

	appleScriptNowPlaying = "spotify_util/now_playing.applescript"
	appleScriptNext       = "spotify_util/playback_next.applescript"
	appleScriptPrev       = "spotify_util/playback_prev.applescript"
	appleScriptToggle     = "spotify_util/playback_toggle.applescript"
	appleScriptRepeat     = "spotify_util/toggle_repeat.applescript"
	appleScriptSuffle     = "spotify_util/toggle_shuffle.applescript"
	appleScriptVolumeUp   = "spotify_util/volume_up.applescript"
	appleScriptVolumeDown = "spotify_util/volume_down.applescript"

	nvimAYUNiSRuntimeVar = "g:ayunis_rtp"
)

type Spotify struct {
	NowPlaying string
	n          *nimvle.Nimvle
	m          *sync.Mutex

	// For initialize GetNowPlaying
	onceRecv   *sync.Once
	initialize *sync.Once
}

func NewSpotify() *Spotify {
	s := &Spotify{
		m:          new(sync.Mutex),
		onceRecv:   new(sync.Once),
		initialize: new(sync.Once),
	}

	return s
}

func (s *Spotify) Init(v *nvim.Nvim, args []string) error {
	s.initialize.Do(func() {
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
		initializedNowPlaying <- struct{}{}
	})

	return nil
}

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	s.onceRecv.Do(func() {
		<-initializedNowPlaying // wait to finish initialize method
	})
	return s.NowPlaying, nil
}

func (s *Spotify) getNowPlaying(rtp string) (string, error) {
	cmd := exec.Command(
		appleScriptCmd, rtp+appleScriptNowPlaying)
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

			time.Sleep(getNowPlayingPollingInterval)
		}
	}()
}

func (s *Spotify) runtimePath(nimvle *nimvle.Nimvle) string {
	irtp, err := nimvle.Eval(nvimAYUNiSRuntimeVar)
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

	return nil
}

func (s *Spotify) Next(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptNext)
}

func (s *Spotify) Prev(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptPrev)
}

func (s *Spotify) Toggle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptToggle)
}

func (s *Spotify) ToggleRepeat(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptRepeat)
}

func (s *Spotify) ToggleShuffle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptSuffle)
}

func (s *Spotify) VolumeUp(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptVolumeUp)
}

func (s *Spotify) VolumeDown(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	return s.exec(v, appleScriptCmd,
		s.runtimePath(nimvle)+appleScriptVolumeDown)
}
