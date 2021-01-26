package variablename

import "errors"

type builder struct {
	operationBuilder OperationBuilder
	operation        Operation
	vr               string
	isPush           bool
}

func createBuilder(operationBuilder OperationBuilder) Builder {
	out := builder{
		operationBuilder: operationBuilder,
		operation:        nil,
		vr:               "",
		isPush:           false,
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

// WithVariableName adds a variableName to the builder
func (app *builder) WithVariableName(vr string) Builder {
	app.vr = vr
	return app
}

// IsPush flags the builder as push
func (app *builder) IsPush() Builder {
	app.isPush = true
	return app
}

// Now builds a new variableName instance
func (app *builder) Now() (VariableName, error) {
	if app.vr == "" {
		return nil, errors.New("the variableName is mandatory in order to build a VariableName instance")
	}

	if app.operation == nil {
		builder := app.operationBuilder.Create()
		if app.isPush {
			builder.IsPush()
		}

		operation, err := builder.Now()
		if err != nil {
			return nil, err
		}

		app.operation = operation
	}

	return createVariableName(app.operation, app.vr), nil
}
