package main

import (
	"github.com/NoahOrberg/AYUNiS.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	s := command.NewSpotify()

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSGetNowPlaying"}, s.GetNowPlaying)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSNext"}, s.Next)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSPrev"}, s.Prev)
		p.HandleFunction(&plugin.FunctionOptions{Name: "InitializeAYUNiS"}, s.Init)
		return nil
	})
}
