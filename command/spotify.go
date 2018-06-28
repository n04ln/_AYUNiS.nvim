package command

import (
	"os/exec"
	"sync"
	"time"

	"github.com/NoahOrberg/nimvle.nvim/nimvle"
	"github.com/neovim/go-client/nvim"
)

type Spotify struct {
	NowPlaying string
	Rtp        string
	m          *sync.Mutex
	retryCount int
}

func (s *Spotify) init() {
	// initialize
	s.m = new(sync.Mutex)
	s.retryCount = 100 // NOTE: about
}

func NewSpotify() *Spotify {
	s := &Spotify{}
	s.init()
	return s
}

func (s *Spotify) Init(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	s.setRuntimePath(nimvle)
	s.pollingNowPlaying()
	return nil
}

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	return s.NowPlaying, nil
}

func (s *Spotify) pollingNowPlaying() {
	rtp := s.Rtp
	go func() {
		for {
			out, err := exec.Command("/usr/bin/osascript", rtp+"spotify_util/now_playing.applescript").Output()
			if err != nil {
				continue
			}

			s.m.Lock()
			s.NowPlaying = string(out[:len(out)-1]) // drop ^@
			s.m.Unlock()

			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (s *Spotify) setRuntimePath(nimvle *nimvle.Nimvle) {
	rtp, err := nimvle.Eval(`dein#get("AYUNiS.nvim")['rtp']`)
	if err != nil {
		panic(err)
	}

	s.Rtp = rtp.(string) + "/"
	return
}

func (s *Spotify) Next(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	go func() {
		// NOTE: not smart...
		for i := 0; i < s.retryCount; i++ {
			if err := nimvle.RedrawStatusLine(); err != nil {
				nimvle.Log(err.Error())
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_next.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}

func (s *Spotify) Prev(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	go func() {
		// NOTE: not smart...
		for i := 0; i < s.retryCount; i++ {
			if err := nimvle.RedrawStatusLine(); err != nil {
				nimvle.Log(err.Error())
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_prev.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}

func (s *Spotify) Toggle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	go func() {
		// NOTE: not smart...
		for i := 0; i < s.retryCount; i++ {
			if err := nimvle.RedrawStatusLine(); err != nil {
				nimvle.Log(err.Error())
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_toggle.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}
