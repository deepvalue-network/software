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
	WithSet(set selectors.Selector) Builder
	MustBeElementEmpty() Builder
	Now() (Transaction, error)
}

// Transaction represents a delete set transaction
type Transaction interface {
	Hash() hash.Hash
	Set() selectors.Selector
	MustBeElementEmpty() bool
}
