package deletes

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures"
	"github.com/deepvalue-network/software/libs/hash"
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
