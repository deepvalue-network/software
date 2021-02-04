package views

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithGenesis(gen spends.Genesis) Builder
	WithSeed(seed string) Builder
	WithAmount(amount uint64) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis view spent
type Genesis interface {
	Hash() hash.Hash
	Genesis() spends.Genesis
	Seed() string
	Amount() uint64
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	ListByGenesis(gen spends.Genesis) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// Service represents a genesis service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
