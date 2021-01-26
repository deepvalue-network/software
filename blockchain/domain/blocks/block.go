package blocks

import (
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hashtree"
)

type block struct {
	tree   hashtree.HashTree
	hashes []hash.Hash
}

func createBlock(
	tree hashtree.HashTree,
	hashes []hash.Hash,
) Block {
	out := block{
		tree:   tree,
		hashes: hashes,
	}

	return &out
}

// Tree returns the block hashtree
func (obj *block) Tree() hashtree.HashTree {
	return obj.tree
}

// Hashes returns the hashes
func (obj *block) Hashes() []hash.Hash {
	return obj.hashes
}
