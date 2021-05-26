package parsers

import "errors"

type registryBuilder struct {
	fetch      FetchRegistry
	register   Register
	unregister Unregister
}

func createRegistryBuilder() RegistryBuilder {
	out := registryBuilder{
		fetch:      nil,
		register:   nil,
		unregister: nil,
	}

	return &out
}

// Create initializes the builder
func (app *registryBuilder) Create() RegistryBuilder {
	return createRegistryBuilder()
}

// WithFetch adds a fetch to the builder
func (app *registryBuilder) WithFetch(fetch FetchRegistry) RegistryBuilder {
	app.fetch = fetch
	return app
}

// WithRegister adds a register to the builder
func (app *registryBuilder) WithRegister(register Register) RegistryBuilder {
	app.register = register
	return app
}

// WithUnregister adds a unregister to the builder
func (app *registryBuilder) WithUnregister(unregister Unregister) RegistryBuilder {
	app.unregister = unregister
	return app
}

// Now builds a new Registry instance
func (app *registryBuilder) Now() (Registry, error) {
	if app.fetch != nil {
		return createRegistryWithFetch(app.fetch), nil
	}

	if app.register != nil {
		return createRegistryWithRegister(app.register), nil
	}

	if app.unregister != nil {
		return createRegistryWithUnregister(app.unregister), nil
	}

	return nil, errors.New("the Registry is invalid")
}
