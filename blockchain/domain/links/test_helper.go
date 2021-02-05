package links

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
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
