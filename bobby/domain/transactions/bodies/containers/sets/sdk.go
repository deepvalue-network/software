package sets

import (
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/sets/deletes"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/sets/saves"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithSave(save saves.Transaction) Builder
	WithDelete(del deletes.Transaction) Builder
	Now() (Transaction, error)
}

// Transaction represents a set container transaction
type Transaction interface {
	Hash() hash.Hash
	IsSave() bool
	Save() saves.Transaction
	IsDelete() bool
	Delete() deletes.Transaction
}
