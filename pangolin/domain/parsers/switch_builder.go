package parsers

import "errors"

type switchBuilder struct {
	variable string
}

func createSwitchBuilder() SwitchBuilder {
	out := switchBuilder{
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *switchBuilder) Create() SwitchBuilder {
	return createSwitchBuilder()
}

// WithVariable adds a avriable to the builder
func (app *switchBuilder) WithVariable(variable string) SwitchBuilder {
	app.variable = variable
	return app
}

// Now builds a new Switch instance
func (app *switchBuilder) Now() (Switch, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Switch instance")
	}

	return createSwitch(app.variable), nil
}
