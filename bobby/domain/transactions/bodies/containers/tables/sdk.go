package tables

import (
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables/deletes"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables/saves"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the transaction builder
type Builder interface {
	Create() Builder
	WithSave(save saves.Transaction) Builder
	WithDelete(del deletes.Transaction) Builder
	Now() (Transaction, error)
}

// Transaction represents a table container transaction
type Transaction interface {
	Hash() hash.Hash
	IsSave() bool
	Save() saves.Transaction
	IsDelete() bool
	Delete() deletes.Transaction
}
