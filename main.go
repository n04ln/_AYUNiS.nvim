package main

import (
	"time"

	"github.com/NoahOrberg/AYUNiS.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	s := command.NewSpotify()

	// 曲名取得
	go func() {
		for i := uint64(0); ; i++ {
			ss := []string{"a", "b"}
			s.NowPlaying = ss[i%uint64(len(ss))]
			time.Sleep(1 * time.Second)
		}
	}()

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "GetNowPlaying"}, s.GetNowPlaying)
		return nil
	})
}
