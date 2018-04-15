package main

import (
	"github.com/NoahOrberg/AYUNiS.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	s := command.NewSpotify()

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "GetNowPlaying"}, s.GetNowPlaying)
		p.HandleFunction(&plugin.FunctionOptions{Name: "InitializeAYUNiS"}, s.Init)
		return nil
	})
}
