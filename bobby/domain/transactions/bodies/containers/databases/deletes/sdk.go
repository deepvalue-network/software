package deletes

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a delete database transaction builder
type Builder interface {
	Create() Builder
	WithDatabase(db selectors.Selector) Builder
	MustBeTableEmpty() Builder
	MustBeSetEmpty() Builder
	Now() (Transaction, error)
}

// Transaction represents a delete database transaction
type Transaction interface {
	Hash() hash.Hash
	Database() selectors.Selector
	MustBeTableEmpty() bool
	MustBeSetEmpty() bool
}
