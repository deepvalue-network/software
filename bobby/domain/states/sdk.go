package states

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/states/overviews"
	"github.com/steve-care-software/products/bobby/domain/transactions"
	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithTransactionsList(transactionsList []transactions.Transactions) Builder
	WithPrevious(prev hash.Hash) Builder
	Now() (State, error)
}

// State represents a database state
type State interface {
	Resource() resources.Immutable
	Transactions() []transactions.Transactions
	Block() blocks.Block
	HasPrevious() bool
	Previous() *hash.Hash
}

// Repository represents the state repository
type Repository interface {
	Last() (State, error)
	Retrieve(hash hash.Hash) (State, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithChain(chain chains.Chain) ServiceBuilder
	Now() (Service, error)
}

// Service represents the state service
type Service interface {
	Prepare(state State) ([]overviews.Overview, error)
	Save(hash hash.Hash) error
}
