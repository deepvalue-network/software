package parsers

import "errors"

type scriptCommandBuilder struct {
	variable string
	values   []ScriptValue
}

func createScriptCommandBuilder() ScriptCommandBuilder {
	out := scriptCommandBuilder{
		variable: "",
		values:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptCommandBuilder) Create() ScriptCommandBuilder {
	return createScriptCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *scriptCommandBuilder) WithVariable(variable string) ScriptCommandBuilder {
	app.variable = variable
	return app
}

// WithValues add values to the builder
func (app *scriptCommandBuilder) WithValues(values []ScriptValue) ScriptCommandBuilder {
	app.values = values
	return app
}

// Now builds a new ScriptValue instance
func (app *scriptCommandBuilder) Now() (ScriptCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a ScriptValue instance")
	}

	if app.values != nil && len(app.values) <= 0 {
		app.values = nil
	}

	if app.values == nil {
		return nil, errors.New("there must be at least 1 ScriptValue instance in order to build a ScriptCommand instance")
	}

	return createScriptCommand(app.variable, app.values), nil
}
