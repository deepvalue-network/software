package deletes

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithTable(table selectors.Selector) Builder
	MustBeRowEmpty() Builder
	Now() (Transaction, error)
}

// Transaction represents a delete table transaction
type Transaction interface {
	Hash() hash.Hash
	Table() selectors.Selector
	MustBeRowEmpty() bool
}
