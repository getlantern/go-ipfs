package iface

import (
	"context"
	"errors"
	"io"
	"time"

	options "github.com/ipfs/go-ipfs/core/coreapi/interface/options"

	cid "gx/ipfs/QmNp85zy9RLrQ5oQD4hPyS39ezrrXpcaa7R4Y9kxdWQLLQ/go-cid"
	ipld "gx/ipfs/QmPN7cwmpcc4DWXb4KTB9dNAJgjuPY69h3npsMfhRrQL9c/go-ipld-format"
)

type Path interface {
	String() string
	Cid() *cid.Cid
	Root() *cid.Cid
	Resolved() bool
}

// TODO: should we really copy these?
//       if we didn't, godoc would generate nice links straight to go-ipld-format
type Node ipld.Node
type Link ipld.Link

type IpnsEntry struct {
	Name  string
	Value Path
}

type Reader interface {
	io.ReadSeeker
	io.Closer
}

type CoreAPI interface {
	Unixfs() UnixfsAPI
	Name() NameAPI
	Key() KeyAPI

	ResolvePath(context.Context, Path) (Path, error)
	ResolveNode(context.Context, Path) (Node, error)
}

type UnixfsAPI interface {
	Add(context.Context, io.Reader) (Path, error)
	Cat(context.Context, Path) (Reader, error)
	Ls(context.Context, Path) ([]*Link, error)
}

type NameAPI interface {
	Publish(ctx context.Context, path Path, opts ...options.NamePublishOption) (*IpnsEntry, error)
	WithValidTime(validTime time.Duration) options.NamePublishOption
	WithKey(key string) options.NamePublishOption

	Resolve(ctx context.Context, name string, opts ...options.NameResolveOption) (Path, error)
	WithRecursive(recursive bool) options.NameResolveOption
	WithLocal(local bool) options.NameResolveOption
	WithNoCache(nocache bool) options.NameResolveOption
}

type KeyAPI interface {
	Generate(ctx context.Context, name string, opts ...options.KeyGenerateOption) (string, error)
	WithAlgorithm(algorithm string) options.KeyGenerateOption
	WithSize(size int) options.KeyGenerateOption

	Rename(ctx context.Context, oldName string, newName string, opts ...options.KeyRenameOption) (string, bool, error)
	WithForce(force bool) options.KeyRenameOption

	List(ctx context.Context) (map[string]string, error) //TODO: better key type?
	Remove(ctx context.Context, name string) (string, error)
}

// type ObjectAPI interface {
// 	New() (cid.Cid, Object)
// 	Get(string) (Object, error)
// 	Links(string) ([]*Link, error)
// 	Data(string) (Reader, error)
// 	Stat(string) (ObjectStat, error)
// 	Put(Object) (cid.Cid, error)
// 	SetData(string, Reader) (cid.Cid, error)
// 	AppendData(string, Data) (cid.Cid, error)
// 	AddLink(string, string, string) (cid.Cid, error)
// 	RmLink(string, string) (cid.Cid, error)
// }

// type ObjectStat struct {
// 	Cid            cid.Cid
// 	NumLinks       int
// 	BlockSize      int
// 	LinksSize      int
// 	DataSize       int
// 	CumulativeSize int
// }

var ErrIsDir = errors.New("object is a directory")
var ErrOffline = errors.New("can't resolve, ipfs node is offline")
