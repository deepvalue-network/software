package owners

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/diamonds/domain/owners"
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
	WithOwner(owner owners.Owner) Builder
	WithGenesis(genesis views.Genesis) Builder
	Now() (Genesis, error)
}

// Genesis represents an owned genesis spent
type Genesis interface {
	Hash() hash.Hash
	Owner() owners.Owner
	Genesis() views.Genesis
}

// Repository represents a genesis repository
type Repository interface {
	List(owner owners.Owner) []hash.Hash
	Retrieve(owner owners.Owner, hash hash.Hash) (Genesis, error)
}

// Service represents a bill service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
