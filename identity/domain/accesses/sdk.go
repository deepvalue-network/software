package accesses

import "github.com/steve-care-software/products/identity/domain/accesses/access"

// NewFactory creates a new factory instance
func NewFactory(encPkBitrate int) Factory {
	accessFactory := access.NewFactory(encPkBitrate)
	return createFactory(accessFactory)
}

// Factory represents an accesses factory
type Factory interface {
	Create() Accesses
}

// Accesses represents accesses
type Accesses interface {
	All() map[string]access.Access
	Create(name string) error
	Fetch(name string) (access.Access, error)
	Delete(name string) error
}
