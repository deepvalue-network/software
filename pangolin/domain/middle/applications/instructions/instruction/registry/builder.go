package registry

import "errors"

type builder struct {
	fetch      Fetch
	register   Register
	unregister string
}

func createBuilder() Builder {
	out := builder{
		fetch:      nil,
		register:   nil,
		unregister: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithFetch adds a fetch to the builder
func (app *builder) WithFetch(fetch Fetch) Builder {
	app.fetch = fetch
	return app
}

// WithRegister adds a register to the builder
func (app *builder) WithRegister(reg Register) Builder {
	app.register = reg
	return app
}

// WithUnregister adds an unregister to the builder
func (app *builder) WithUnregister(unregister string) Builder {
	app.unregister = unregister
	return app
}

// Now builds a new Registry instance
func (app *builder) Now() (Registry, error) {
	if app.fetch != nil {
		return createRegistryWithFetch(app.fetch), nil
	}

	if app.register != nil {
		return createRegistryWithRegister(app.register), nil
	}

	if app.unregister != "" {
		return createRegistryWithUnregister(app.unregister), nil
	}

	return nil, errors.New("the Registry is invalid")
}
