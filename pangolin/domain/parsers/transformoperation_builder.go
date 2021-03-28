package parsers

import "errors"

type transformOperationBuilder struct {
	input  string
	result string
}

func createTransformOperationBuilder() TransformOperationBuilder {
	out := transformOperationBuilder{
		input:  "",
		result: "",
	}

	return &out
}

// Create initializes the builder
func (app *transformOperationBuilder) Create() TransformOperationBuilder {
	return createTransformOperationBuilder()
}

// WithInput adds the input to the builder
func (app *transformOperationBuilder) WithInput(input string) TransformOperationBuilder {
	app.input = input
	return app
}

// WithResult adds the result to the builder
func (app *transformOperationBuilder) WithResult(result string) TransformOperationBuilder {
	app.result = result
	return app
}

// Now builds a new TransformOperation instance
func (app *transformOperationBuilder) Now() (TransformOperation, error) {
	if app.input == "" {
		return nil, errors.New("the input string is mandatory in order to build a TransformOperation instance")
	}

	if app.result == "" {
		return nil, errors.New("the result string is mandatory in order to build a TransformOperation instance")
	}

	return createTransformOperation(app.input, app.result), nil
}
