package containers

import (
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/databases"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/graphbases"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/sets"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithGraphbase(graph graphbases.Transaction) Builder
	WithDatabase(db databases.Transaction) Builder
	WithTable(table tables.Transaction) Builder
	WithSet(set sets.Transaction) Builder
	Now() (Transaction, error)
}

// Transaction represents a container transaction
type Transaction interface {
	Hash() hash.Hash
	IsGraphbase() bool
	Graphbase() graphbases.Transaction
	IsDatabase() bool
	Database() databases.Transaction
	IsTable() bool
	Table() tables.Transaction
	IsSet() bool
	Set() sets.Transaction
}
