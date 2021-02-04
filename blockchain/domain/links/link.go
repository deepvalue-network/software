package links

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hash"
)

type link struct {
	hash          hash.Hash    `hydro:"Hash, Hash"`
	index         uint         `hydro:"Index, Index"`
	prevMinedLink hash.Hash    `hydro:"PrevMinedLink, PrevMinedLink"`
	nextBlock     blocks.Block `hydro:"NextBlock, NextBlock"`
}

func createLink(
	hash hash.Hash,
	index uint,
	prevMinedLink hash.Hash,
	nextBlock blocks.Block,
) Link {
	out := link{
		hash:          hash,
		index:         index,
		prevMinedLink: prevMinedLink,
		nextBlock:     nextBlock,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *link) Index() uint {
	return obj.index
}

// PrevMinedLink returns the previous mined link hash
func (obj *link) PrevMinedLink() hash.Hash {
	return obj.prevMinedLink
}

// NextBlock returns the next block
func (obj *link) NextBlock() blocks.Block {
	return obj.nextBlock
}
