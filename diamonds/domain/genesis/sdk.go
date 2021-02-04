package genesis

import (
	"time"

	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	ActiveOn(activeOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis diamond
type Genesis interface {
	Hash() hash.Hash
	Owner() []hash.Hash
	CreatedOn() time.Time
	ActiveOn() time.Time
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// Service represents a genesis service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
