package parsers

import "errors"

type variableNameBuilder struct {
	local string
}

func createVariableNameBuilder() VariableNameBuilder {
	out := variableNameBuilder{
		local: "",
	}

	return &out
}

// Create initializes the builder
func (app *variableNameBuilder) Create() VariableNameBuilder {
	return createVariableNameBuilder()
}

// WithLocal adds a local variable to the builder
func (app *variableNameBuilder) WithLocal(local string) VariableNameBuilder {
	app.local = local
	return app
}

// Now builds a new VariableName instance
func (app *variableNameBuilder) Now() (VariableName, error) {
	if app.local != "" {
		return createVariableNameWithLocal(app.local), nil
	}

	return nil, errors.New("the VariableName is invalid")
}
