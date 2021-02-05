package saves

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets/schemas"
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
	WithDatabase(db selectors.Selector) Builder
	WithSchema(schema schemas.Schema) Builder
	WithSet(set selectors.Selector) Builder
	Now() (Transaction, error)
}

// Transaction represents a set container transaction
type Transaction interface {
	Hash() hash.Hash
	Database() selectors.Selector
	Schema() schemas.Schema
	HasSet() bool
	Set() selectors.Selector
}
