package parsers

import "errors"

type unregisterBuilder struct {
	variable string
}

func createUnregisterBuilder() UnregisterBuilder {
	out := unregisterBuilder{
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *unregisterBuilder) Create() UnregisterBuilder {
	return createUnregisterBuilder()
}

// WithVariable adds a variable to the builder
func (app *unregisterBuilder) WithVariable(name string) UnregisterBuilder {
	app.variable = name
	return app
}

// Now builds a new Unregister instance
func (app *unregisterBuilder) Now() (Unregister, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Unregister instance")
	}

	return createUnregister(app.variable), nil
}
