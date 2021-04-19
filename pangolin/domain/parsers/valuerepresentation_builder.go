package parsers

import "errors"

type valueRepresentationBuilder struct {
	value    Value
	variable string
}

func createValueRepresentationBuilder() ValueRepresentationBuilder {
	out := valueRepresentationBuilder{
		value:    nil,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *valueRepresentationBuilder) Create() ValueRepresentationBuilder {
	return createValueRepresentationBuilder()
}

// WithValue adds a value to the builder
func (app *valueRepresentationBuilder) WithValue(value Value) ValueRepresentationBuilder {
	app.value = value
	return app
}

// WithVariable adds a variable to the builder
func (app *valueRepresentationBuilder) WithVariable(variable string) ValueRepresentationBuilder {
	app.variable = variable
	return app
}

// Now builds a new ValueRepresentation instance
func (app *valueRepresentationBuilder) Now() (ValueRepresentation, error) {
	if app.value != nil {
		return createValueRepresentationWithValue(app.value), nil
	}

	if app.variable != "" {
		return createValueRepresentationWithVariable(app.variable), nil
	}

	return nil, errors.New("the ValueRepresentation is invalid")
}
