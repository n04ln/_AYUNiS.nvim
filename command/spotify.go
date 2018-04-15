package command

import "github.com/neovim/go-client/nvim"

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

func (s *Spotify) GetNowPlaying(v *nvim.Nvim, args []string) (string, error) {
	return s.NowPlaying, nil
}
