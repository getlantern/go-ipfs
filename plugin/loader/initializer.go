package loader

import (
	"github.com/getlantern/go-ipfs/core/coredag"
	"github.com/getlantern/go-ipfs/plugin"

	format "github.com/ipfs/go-ipld-format"
)

func initialize(plugins []plugin.Plugin) error {
	for _, p := range plugins {
		err := p.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func run(plugins []plugin.Plugin) error {
	for _, pl := range plugins {
		err := runIPLDPlugin(pl)
		if err != nil {
			return err
		}
	}
	return nil
}

func runIPLDPlugin(pl plugin.Plugin) error {
	ipldpl, ok := pl.(plugin.PluginIPLD)
	if !ok {
		return nil
	}

	err := ipldpl.RegisterBlockDecoders(format.DefaultBlockDecoder)
	if err != nil {
		return err
	}

	return ipldpl.RegisterInputEncParsers(coredag.DefaultInputEncParsers)
}
