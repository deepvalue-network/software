package overviews

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/structures"
	"github.com/deepvalue-network/software/bobby/domain/transactions"
)

type validTransaction struct {
	trx        transactions.Transaction
	structures []structures.Structure
	chain      chains.Chain
	block      blocks.Block
}

func createValidTransaction(
	trx transactions.Transaction,
	structures []structures.Structure,
	chain chains.Chain,
	block blocks.Block,
) ValidTransaction {
	out := validTransaction{
		trx:        trx,
		structures: structures,
		chain:      chain,
		block:      block,
	}

	return &out
}

// Transaction returns the transaction
func (obj *validTransaction) Transaction() transactions.Transaction {
	return obj.trx
}

// Structures returns the structures
func (obj *validTransaction) Structures() []structures.Structure {
	return obj.structures
}

// Chain returns the chain
func (obj *validTransaction) Chain() chains.Chain {
	return obj.chain
}

// Block returns the block
func (obj *validTransaction) Block() blocks.Block {
	return obj.block
}
