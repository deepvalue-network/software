package deletes

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Processor represents a transaction processor
type Processor interface {
	Execute(trx Transaction) ([]structures.Structure, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithGraphbase(graphbase selectors.Selector) Builder
	MustBeGraphbaseEmpty() Builder
	Now() (Transaction, error)
}

// Transaction represents a delete graphbase transaction
type Transaction interface {
	Hash() hash.Hash
	Graphbase() selectors.Selector
	MustBeGraphbaseEmpty() bool
}
