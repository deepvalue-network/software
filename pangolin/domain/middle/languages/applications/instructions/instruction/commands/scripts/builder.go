package scripts

import "errors"

type builder struct {
	variable string
	values   []Value
}

func createBuilder() Builder {
	out := builder{
		variable: "",
		values:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithValues add values to the builder
func (app *builder) WithValues(values []Value) Builder {
	app.values = values
	return app
}

// Now builds a new Script instance
func (app *builder) Now() (Script, error) {
	if app.variable != "" {
		return nil, errors.New("the variable is mandatory in order to build a Script instance")
	}

	if app.values != nil && len(app.values) <= 0 {
		return nil, errors.New("the values are mandatory in order to build a Script instance")
	}

	return createScript(app.variable, app.values), nil
}
