package coredag

import (
	"io"
	"io/ioutil"

	ipldcbor "github.com/ipfs/go-ipld-cbor"
	node "github.com/ipfs/go-ipld-format"
)

func cborJSONParser(r io.Reader, mhType uint64, mhLen int) ([]node.Node, error) {
	nd, err := ipldcbor.FromJson(r, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []node.Node{nd}, nil
}

func cborRawParser(r io.Reader, mhType uint64, mhLen int) ([]node.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nd, err := ipldcbor.Decode(data, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []node.Node{nd}, nil
}
