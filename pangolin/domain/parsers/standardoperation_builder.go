package parsers

import "errors"

type standardOperationBuilder struct {
	first  Identifier
	second Identifier
	result VariableName
}

func createStandardOperationBuilder() StandardOperationBuilder {
	out := standardOperationBuilder{
		first:  nil,
		second: nil,
		result: nil,
	}

	return &out
}

// Create initializes the builder
func (app *standardOperationBuilder) Create() StandardOperationBuilder {
	return createStandardOperationBuilder()
}

// WithFirst adds the first variable to the builder
func (app *standardOperationBuilder) WithFirst(first Identifier) StandardOperationBuilder {
	app.first = first
	return app
}

// WithSecond adds the second variable to the builder
func (app *standardOperationBuilder) WithSecond(second Identifier) StandardOperationBuilder {
	app.second = second
	return app
}

// WithResult adds the result variable to the builder
func (app *standardOperationBuilder) WithResult(result VariableName) StandardOperationBuilder {
	app.result = result
	return app
}

// Now builds a new StandardOperation instance
func (app *standardOperationBuilder) Now() (StandardOperation, error) {
	if app.first == nil {
		return nil, errors.New("the first variable is mandatory in order to build a StandardOperation instance")
	}

	if app.second == nil {
		return nil, errors.New("the second variable is mandatory in order to build a StandardOperation instance")
	}

	if app.result == nil {
		return nil, errors.New("the result variable is mandatory in order to build a StandardOperation instance")
	}

	return createStandardOperation(app.first, app.second, app.result), nil
}
