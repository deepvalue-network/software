package overviews

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/bobby/domain/transactions"
    "errors"
)

type validTransactionBuilder struct {
	trx        transactions.Transaction
	structures []structures.Structure
	chain      chains.Chain
	block      blocks.Block
}

func createValidTransactionBuilder() ValidTransactionBuilder {
    out := validTransactionBuilder{
        trx: nil,
        structures: nil,
        chain: nil,
        block: nil,
    }

    return &out
}

// Create initializes the builder
func (app *validTransactionBuilder) Create() ValidTransactionBuilder {
    return createValidTransactionBuilder()
}

// WithTransaction adds a transaction to the builder
func (app *validTransactionBuilder) WithTransaction(trx transactions.Transaction) ValidTransactionBuilder {
    app.trx = trx
    return app
}

// WithStructures add structures to the builder
func (app *validTransactionBuilder) WithStructures(structures []structures.Structure) ValidTransactionBuilder {
    app.structures = structures
    return app
}

// WithChain add a chain to the builder
func (app *validTransactionBuilder) WithChain(chain chains.Chain) ValidTransactionBuilder {
    app.chain = chain
    return app
}

// WithBlock add a block to the builder
func (app *validTransactionBuilder) WithBlock(block blocks.Block) ValidTransactionBuilder {
    app.block = block
    return app
}

// Now builds a new ValidTransaction instance
func (app *validTransactionBuilder) Now() (ValidTransaction, error) {
    if app.trx == nil {
        return nil, errors.New("the transaction is mandatory in order to build a ValidTransaction instance")
    }

    if app.structures != nil && len(app.structures) <= 0 {
        app.structures = nil
    }

    if app.structures == nil {
        return nil, errors.New("there must be at least 1 structure instance in order to build a ValidTransaction instance")
    }

    if app.chain == nil {
        return nil, errors.New("the chain is mandatory in order to build a ValidTransaction instance")
    }

    if app.block == nil {
        return nil, errors.New("the block is mandatory in order to build a ValidTransaction instance")
    }

    return createValidTransaction(app.trx, app.structures, app.chain, app.block), nil
}
