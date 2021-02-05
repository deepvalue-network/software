package users

import "github.com/deepvalue-network/software/identity/domain/accesses"

// NewBuilder creates a new builder instance
func NewBuilder(encPkBitrate int) Builder {
	accessesFactory := accesses.NewFactory(encPkBitrate)
	return createBuilder(accessesFactory)
}

// NewPointer returns a new user pointer
func NewPointer() interface{} {
	return new(user)
}

// Builder represents a user builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithAccesses(accesses accesses.Accesses) Builder
	Now() (User, error)
}

// User represents a user
type User interface {
	Name() string
	Seed() string
	Accesses() accesses.Accesses
}

// Repository represents a user repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (User, error)
}

// Service represents a user service
type Service interface {
	Insert(user User, password string) error
	Update(original User, updated User, originalPass string) error
	UpdateWithPassword(original User, updated User, originalPass string, updatedPassword string) error
	Delete(user User, password string) error
}
