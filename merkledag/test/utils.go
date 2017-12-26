package mdutils

import (
	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/getlantern/go-ipfs/blocks/blockstore"
	bsrv "github.com/getlantern/go-ipfs/blockservice"
	"github.com/getlantern/go-ipfs/exchange/offline"
	dag "github.com/getlantern/go-ipfs/merkledag"
)

func Mock() dag.DAGService {
	return dag.NewDAGService(Bserv())
}

func Bserv() bsrv.BlockService {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	return bsrv.New(bstore, offline.Exchange(bstore))
}
