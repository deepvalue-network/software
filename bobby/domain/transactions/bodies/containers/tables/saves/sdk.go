package saves

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithDatabase(db selectors.Selector) Builder
	WithSchema(schema schemas.Schema) Builder
	WithTable(table selectors.Selector) Builder
	Now() (Transaction, error)
}

// Transaction represents a table container transaction
type Transaction interface {
	Hash() hash.Hash
	Database() selectors.Selector
	Schema() schemas.Schema
	HasTable() bool
	Table() selectors.Selector
}
