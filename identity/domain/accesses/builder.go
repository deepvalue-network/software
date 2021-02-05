package accesses

import "github.com/deepvalue-network/software/identity/domain/accesses/access"

type builder struct {
	accessFactory access.Factory
	mp            map[string]access.Access
}

func createBuilder(
	accessFactory access.Factory,
) Builder {
	out := builder{
		accessFactory: accessFactory,
		mp:            nil,
	}

	return &out
}

// Create initialzies the builder
func (app *builder) Create() Builder {
	return createBuilder(app.accessFactory)
}

// WithMap adds a map to the builder
func (app *builder) WithMap(mp map[string]access.Access) Builder {
	app.mp = mp
	return app
}

// Now builds a new Accesses instance
func (app *builder) Now() (Accesses, error) {
	if app.mp == nil {
		app.mp = map[string]access.Access{}
	}

	return createAccesses(app.accessFactory, app.mp), nil
}
