package saves

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
	WithGraphbase(graphbase selectors.Selector) Builder
	WithName(name string) Builder
	WithDatabase(db selectors.Selector) Builder
	Now() (Transaction, error)
}

// Transaction represents a database create transaction
type Transaction interface {
	Hash() hash.Hash
	Graphbase() selectors.Selector
	Name() string
	HasDatabase() bool
	Database() selectors.Selector
}
