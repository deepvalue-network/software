package parsers

import "errors"

type variableNameBuilder struct {
	global string
	local  string
}

func createVariableNameBuilder() VariableNameBuilder {
	out := variableNameBuilder{
		global: "",
		local:  "",
	}

	return &out
}

// Create initializes the builder
func (app *variableNameBuilder) Create() VariableNameBuilder {
	return createVariableNameBuilder()
}

// WithGlobal adds a global variable to the builder
func (app *variableNameBuilder) WithGlobal(global string) VariableNameBuilder {
	app.global = global
	return app
}

// WithLocal adds a local variable to the builder
func (app *variableNameBuilder) WithLocal(local string) VariableNameBuilder {
	app.local = local
	return app
}

// Now builds a new VariableName instance
func (app *variableNameBuilder) Now() (VariableName, error) {
	if app.global != "" {
		return createVariableNameWithGlobal(app.global), nil
	}

	if app.local != "" {
		return createVariableNameWithLocal(app.local), nil
	}

	return nil, errors.New("the VariableName is invalid")
}
