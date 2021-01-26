package parsers

import "errors"

type identifierBuilder struct {
	variableName VariableName
	constant     string
}

func createIdentifierBuilder() IdentifierBuilder {
	out := identifierBuilder{
		variableName: nil,
		constant:     "",
	}

	return &out
}

// Create initializes the builder
func (app *identifierBuilder) Create() IdentifierBuilder {
	return createIdentifierBuilder()
}

// WithVariable adds a variable to the builder
func (app *identifierBuilder) WithVariable(variable VariableName) IdentifierBuilder {
	app.variableName = variable
	return app
}

// WithConstant adds a constant to the builder
func (app *identifierBuilder) WithConstant(constant string) IdentifierBuilder {
	app.constant = constant
	return app
}

// Now builds a new Identifier instance
func (app *identifierBuilder) Now() (Identifier, error) {
	if app.variableName != nil {
		return createIdentifierWithVariableName(app.variableName), nil
	}

	if app.constant != "" {
		return createIdentifierWithConstant(app.constant), nil
	}

	return nil, errors.New("the Identifier is invalid")
}
