package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
)

// CreateBlockForTests creates a block for tests
func CreateBlockForTests() Block {
	block := blocks.CreateBlockForTests()
	results := "some results"
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithBlock(block).WithResults(results).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
