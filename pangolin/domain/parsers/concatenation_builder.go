package parsers

import "errors"

type concatenationBuilder struct {
	operation StandardOperation
}

func createConcatenationBuilder() ConcatenationBuilder {
	out := concatenationBuilder{
		operation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *concatenationBuilder) Create() ConcatenationBuilder {
	return createConcatenationBuilder()
}

// WithOperation add a standard operation to the builder
func (app *concatenationBuilder) WithOperation(operation StandardOperation) ConcatenationBuilder {
	app.operation = operation
	return app
}

// Now builds a new Concatenation instance
func (app *concatenationBuilder) Now() (Concatenation, error) {
	if app.operation == nil {
		return nil, errors.New("the StandardOperation is mandatory in order to build a Concatenation instance")
	}

	return createConcatenation(app.operation), nil
}
