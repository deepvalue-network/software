package blocks

import "github.com/deepvalue-network/software/libs/hash"

// CreateBlockForTests creates a new block instance for tests
func CreateBlockForTests() Block {
	hashAdapter := hash.NewAdapter()
	firstHash, err := hashAdapter.FromBytes([]byte("first hash"))
	if err != nil {
		panic(err)
	}

	secondHash, err := hashAdapter.FromBytes([]byte("second hash"))
	if err != nil {
		panic(err)
	}

	thirdHash, err := hashAdapter.FromBytes([]byte("third hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHashes([]hash.Hash{
		*firstHash,
		*secondHash,
		*thirdHash,
	}).Now()

	if err != nil {
		panic(err)
	}

	return ins
}
