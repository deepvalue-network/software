package transform

import "errors"

type builder struct {
	operationBuilder OperationBuilder
	operation        Operation
	result           string
	input            string
	isPop            bool
}

func createBuilder(operationBuilder OperationBuilder) Builder {
	out := builder{
		operationBuilder: operationBuilder,
		operation:        nil,
		result:           "",
		input:            "",
		isPop:            false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.operationBuilder)
}

// WithOperation adds an operation to the builder
func (app *builder) WithOperation(operation Operation) Builder {
	app.operation = operation
	return app
}

// WithResult adds a result to the builder
func (app *builder) WithResult(result string) Builder {
	app.result = result
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// IsPop flags the builder as pop
func (app *builder) IsPop() Builder {
	app.isPop = true
	return app
}

// Now builds a new Transform instance
func (app *builder) Now() (Transform, error) {
	if app.result == "" {
		return nil, errors.New("the result is mandatory in order to build a Transform instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Transform instance")
	}

	if app.operation == nil {
		builder := app.operationBuilder.Create()
		if app.isPop {
			builder.IsPop()
		}

		operation, err := builder.Now()
		if err != nil {
			return nil, err
		}

		app.operation = operation
	}

	return createTransform(app.operation, app.result, app.input), nil
}
