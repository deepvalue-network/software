package hydro

import "errors"

type adapterBuilder struct {
	manager Manager
}

func createAdapterBuilder() AdapterBuilder {
	out := adapterBuilder{
		manager: nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder()
}

// WithManager adds a manager to the builder
func (app *adapterBuilder) WithManager(manager Manager) AdapterBuilder {
	app.manager = manager
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.manager == nil {
		return nil, errors.New("the manager is mandatory in order to build an Adapter instance")
	}

	return createAdapter(app.manager), nil
}
