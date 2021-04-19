package parsers

import "errors"

type headCommandBuilder struct {
	variable string
	values   []HeadValue
}

func createHeadCommandBuilder() HeadCommandBuilder {
	out := headCommandBuilder{
		variable: "",
		values:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *headCommandBuilder) Create() HeadCommandBuilder {
	return createHeadCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *headCommandBuilder) WithVariable(variable string) HeadCommandBuilder {
	app.variable = variable
	return app
}

// WithValues add values to the builder
func (app *headCommandBuilder) WithValues(values []HeadValue) HeadCommandBuilder {
	app.values = values
	return app
}

// Now builds a new HeadCommand instance
func (app *headCommandBuilder) Now() (HeadCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an HeadCommand instance")
	}

	if app.values != nil && len(app.values) <= 0 {
		app.values = nil
	}

	if app.values == nil {
		return nil, errors.New("there must be at least 1 HeadValue instance in order to build an HeadCommand instance")
	}

	return createHeadCommand(app.variable, app.values), nil
}
