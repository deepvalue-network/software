package parsers

import "errors"

type transformOperationBuilder struct {
	input  Identifier
	result VariableName
}

func createTransformOperationBuilder() TransformOperationBuilder {
	out := transformOperationBuilder{
		input:  nil,
		result: nil,
	}

	return &out
}

// Create initializes the builder
func (app *transformOperationBuilder) Create() TransformOperationBuilder {
	return createTransformOperationBuilder()
}

// WithInput adds the input to the builder
func (app *transformOperationBuilder) WithInput(input Identifier) TransformOperationBuilder {
	app.input = input
	return app
}

// WithResult adds the result to the builder
func (app *transformOperationBuilder) WithResult(result VariableName) TransformOperationBuilder {
	app.result = result
	return app
}

// Now builds a new TransformOperation instance
func (app *transformOperationBuilder) Now() (TransformOperation, error) {
	if app.input == nil {
		return nil, errors.New("the input Identifier is mandatory in order to build a TransformOperation instance")
	}

	if app.result == nil {
		return nil, errors.New("the result VariableName is mandatory in order to build a TransformOperation instance")
	}

	return createTransformOperation(app.input, app.result), nil
}
