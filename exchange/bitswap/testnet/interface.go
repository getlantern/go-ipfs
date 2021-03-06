package bitswap

import (
	bsnet "github.com/getlantern/go-ipfs/exchange/bitswap/network"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-testutil"
)

type Network interface {
	Adapter(testutil.Identity) bsnet.BitSwapNetwork

	HasPeer(peer.ID) bool
}
