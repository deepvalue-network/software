package links

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hash"
)

// CreateLinkForTests creates link instance for tests
func CreateLinkForTests() Link {
	prevLinkHash, err := hash.NewAdapter().FromBytes([]byte("prev link hash"))
	if err != nil {
		panic(err)
	}

	index := uint(45)
	next := blocks.CreateBlockForTests()
	ins, err := NewBuilder().Create().WithPreviousMinedLink(*prevLinkHash).WithNextBlock(next).WithIndex(index).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
