package registry

import "errors"

type registerBuilder struct {
	variable string
	index    Index
}

func createRegisterBuilder() RegisterBuilder {
	out := registerBuilder{
		variable: "",
		index:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *registerBuilder) Create() RegisterBuilder {
	return createRegisterBuilder()
}

// WithVariable adds a variable to the builder
func (app *registerBuilder) WithVariable(variable string) RegisterBuilder {
	app.variable = variable
	return app
}

// WithIndex adds an index to the builder
func (app *registerBuilder) WithIndex(index Index) RegisterBuilder {
	app.index = index
	return app
}

// Now builds a new Register instance
func (app *registerBuilder) Now() (Register, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Register instance")
	}

	if app.index != nil {
		return createRegisterWithIndex(app.variable, app.index), nil
	}

	return createRegister(app.variable), nil
}
