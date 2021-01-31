package accesses

import "github.com/steve-care-software/products/identity/domain/accesses/access"

type factory struct {
	accessFactory access.Factory
}

func createFactory(
	accessFactory access.Factory,
) Factory {
	out := factory{
		accessFactory: accessFactory,
	}

	return &out
}

// Create creates an accesses instance
func (app *factory) Create() Accesses {
	mp := map[string]access.Access{}
	return createAccesses(app.accessFactory, mp)
}
