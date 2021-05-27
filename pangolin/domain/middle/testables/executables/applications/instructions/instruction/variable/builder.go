package variable

import (
	"errors"

	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
)

type builder struct {
	name  string
	value var_value.Value
}

func createBuilder() Builder {
	out := builder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(val var_value.Value) Builder {
	app.value = val
	return app
}

// Now builds a new Variable instance
func (app *builder) Now() (Variable, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Variable instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Variable instance")
	}

	return createVariable(
		app.name,
		app.value,
	), nil
}
