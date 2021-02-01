package accesses

import "github.com/steve-care-software/products/identity/domain/accesses/access"

// NewFactory creates a new factory instance
func NewFactory(encPkBitrate int) Factory {
	builder := NewBuilder(encPkBitrate)
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder(encPkBitrate int) Builder {
	accessFactory := access.NewFactory(encPkBitrate)
	return createBuilder(accessFactory)
}

// NewPointer returns a new accesses pointer
func NewPointer() interface{} {
	return new(accesses)
}

// Factory represents an accesses factory
type Factory interface {
	Create() (Accesses, error)
}

// Builder represents an accesses builder
type Builder interface {
	Create() Builder
	WithMap(mp map[string]access.Access) Builder
	Now() (Accesses, error)
}

// Accesses represents accesses
type Accesses interface {
	All() map[string]access.Access
	Create(name string) error
	Fetch(name string) (access.Access, error)
	Delete(name string) error
}
