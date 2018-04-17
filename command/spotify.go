package command

import (
	"os/exec"

	"github.com/NoahOrberg/nimvle.nvim/nimvle"
	"github.com/neovim/go-client/nvim"
)

type Spotify struct {
	NowPlaying string
	Rtp        string
}

func (s *Spotify) init() {
	// initialize
}

func NewSpotify() *Spotify {
	s := &Spotify{}
	s.init()
	return s
}

func (s *Spotify) Init(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	s.setRuntimePath(nimvle)
	s.pollingNowPlaying(nimvle)
	return nil
}

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	return s.NowPlaying, nil
}

func (s *Spotify) pollingNowPlaying(nimvle *nimvle.Nimvle) {
	rtp := s.Rtp
	go func() {
		for {
			out, err := exec.Command("/usr/bin/osascript", rtp+"spotify_util/now_playing.applescript").Output()
			if err != nil {
				nimvle.Log(err.Error())
				continue
			}

			s.NowPlaying = string(out[:len(out)-1]) // drop ^@
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

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_next.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}

func (s *Spotify) Prev(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_prev.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}

func (s *Spotify) Toggle(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")

	_, err := exec.Command("/usr/bin/osascript", s.Rtp+"spotify_util/playback_toggle.applescript").Output()
	if err != nil {
		nimvle.Log(err.Error())
	}

	return nil
}
