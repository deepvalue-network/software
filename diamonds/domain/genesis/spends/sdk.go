package spends

import (
	"time"

	"github.com/steve-care-software/products/diamonds/domain/genesis"
	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithHashedAmount(hashedAmount hash.Hash) Builder
	WithEncryptedSeed(encSeed []byte) Builder
	WithGenesis(gen genesis.Genesis) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a spent genesis
type Genesis interface {
	Hash() hash.Hash
	Amount() []byte
	Seed() []byte
	Genesis() genesis.Genesis
	CreatedOn() time.Time
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	ListByGenesis(gen genesis.Genesis) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// Service represents a genesis service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
