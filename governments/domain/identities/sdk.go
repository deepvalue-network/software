package identities

import "github.com/deepvalue-network/software/governments/domain/identities/shareholders"

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithShareHolders(shareHolders shareholders.ShareHolders) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Seed() string
	ShareHolders() shareholders.ShareHolders
}

// Repository represents an identity repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (Identity, error)
}

// Service represents a shareholders service
type Service interface {
	Insert(ins Identity, password string) error
	Update(origin Identity, updated Identity, password string) error
	UpdateWithPassword(origin Identity, updated Identity, originalPassword string, updatedPassword string) error
	Delete(ins Identity, password string) error
}
