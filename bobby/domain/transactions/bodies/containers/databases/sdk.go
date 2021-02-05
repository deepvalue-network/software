package databases

import (
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/databases/deletes"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/databases/saves"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithDelete(del deletes.Transaction) Builder
	WithSave(save saves.Transaction) Builder
	Now() (Transaction, error)
}

// Transaction represents a database container transaction
type Transaction interface {
	Hash() hash.Hash
	IsDelete() bool
	Delete() deletes.Transaction
	IsSave() bool
	Save() saves.Transaction
}
