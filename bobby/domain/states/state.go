package states

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/transactions"
	"github.com/steve-care-software/products/libs/hash"
)

type state struct {
	resource resources.Immutable
	trxList  []transactions.Transactions
	block    blocks.Block
	prev     *hash.Hash
}

func createState(
	resource resources.Immutable,
	trxList []transactions.Transactions,
	block blocks.Block,
) State {
	return createStateInternally(resource, trxList, block, nil)
}

func createStateWithPrevious(
	resource resources.Immutable,
	trxList []transactions.Transactions,
	block blocks.Block,
	prev *hash.Hash,
) State {
	return createStateInternally(resource, trxList, block, prev)
}

func createStateInternally(
	resource resources.Immutable,
	trxList []transactions.Transactions,
	block blocks.Block,
	prev *hash.Hash,
) State {
	out := state{
		resource: resource,
		trxList:  trxList,
		block:    block,
		prev:     prev,
	}

	return &out
}

// Resource returns the resource
func (obj *state) Resource() resources.Immutable {
	return obj.resource
}

// Transactions returns the transactions list
func (obj *state) Transactions() []transactions.Transactions {
	return obj.trxList
}

// Block returns the block
func (obj *state) Block() blocks.Block {
	return obj.block
}

// HasPrevious returns true if there is a previous hash, false otherwise
func (obj *state) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous hash, if any
func (obj *state) Previous() *hash.Hash {
	return obj.prev
}
