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
