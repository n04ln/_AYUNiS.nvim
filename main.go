package main

import (
	"github.com/n04ln/AYUNiS.nvim/command"

	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	s := command.NewSpotify()

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSGetNowPlaying"}, s.GetNowPlaying)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSNext"}, s.Next)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSPrev"}, s.Prev)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSToggle"}, s.Toggle)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSToggleRepeat"}, s.ToggleRepeat)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSToggleShuffle"}, s.ToggleShuffle)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSVolumeUp"}, s.VolumeUp)
		p.HandleFunction(&plugin.FunctionOptions{Name: "AYUNiSVolumeDown"}, s.VolumeDown)
		p.HandleFunction(&plugin.FunctionOptions{Name: "InitializeAYUNiS"}, s.Init)
		return nil
	})
}
