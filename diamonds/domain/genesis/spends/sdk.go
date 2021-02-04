package spends

import (
	"time"

	domain_genesis "github.com/steve-care-software/products/diamonds/domain/genesis"
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
	WithAmount(amount uint64) Builder
	WithHashedAmount(amount hash.Hash) Builder
	WithSeed(seed string) Builder
	WithHashedSeed(seed hash.Hash) Builder
	WithGenesis(gen domain_genesis.Genesis) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a spent genesis
type Genesis interface {
	Hash() hash.Hash
	Amount() hash.Hash
	Seed() hash.Hash
	Genesis() domain_genesis.Genesis
	CreatedOn() time.Time
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	ListByGenesis(gen domain_genesis.Genesis) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// Service represents a genesis service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
