package accesses

import "github.com/steve-care-software/products/identity/domain/accesses/access"

// Factory represents an accesses factory
type Factory interface {
	Create() (Accesses, error)
}

// Accesses represents accesses
type Accesses interface {
	Create(name string) error
	Fetch(name string) (access.Access, error)
	Delete(name string) (access.Access, error)
}
