package command

import (
	"os/exec"

	"github.com/NoahOrberg/nimvle.nvim/nimvle"
	"github.com/neovim/go-client/nvim"
)

type Spotify struct {
	NowPlaying string
}

func (s *Spotify) init() {
	// TODO: 現在流れている曲取得
	s.NowPlaying = "a"
}

func NewSpotify() *Spotify {
	s := &Spotify{}
	s.init()
	return s
}

func (s *Spotify) Init(v *nvim.Nvim, args []string) error {
	nimvle := nimvle.New(v, "AYUNiS.nvim")
	pollingNowPlaying(nimvle)
	return nil
}

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	return s.NowPlaying, nil
}

func (s *Spotify) pollingNowPlaying(nimvle *nimvle.Nimve) {
	go func() {
		for i := uint64(0); ; i++ {
			out, err := exec.Command("/usr/bin/osascript", "spotify_util/now_playing.applescript").Output()
			if err != nil {
				nimvle.Log(err.Error())
			}

			s.NowPlaying = string(out)
		}
	}()

}
