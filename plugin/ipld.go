package plugin

import (
	"github.com/getlantern/go-ipfs/core/coredag"

	node "github.com/ipfs/go-ipld-format"
)

// PluginIPLD is an interface that can be implemented to add handlers for
// for different IPLD formats
type PluginIPLD interface {
	Plugin

	RegisterBlockDecoders(dec node.BlockDecoder) error
	RegisterInputEncParsers(iec coredag.InputEncParsers) error
}
