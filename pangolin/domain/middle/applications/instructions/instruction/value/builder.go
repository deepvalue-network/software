package value

import (
	"errors"

	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
)

type builder struct {
	operationBuilder OperationBuilder
	val              var_value.Value
	op               Operation
	isPrint          bool
}

func createBuilder(operationBuilder OperationBuilder) Builder {
	out := builder{
		operationBuilder: operationBuilder,
		val:              nil,
		op:               nil,
		isPrint:          false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.operationBuilder)
}

// WithValue adds a value to the builder
func (app *builder) WithValue(val var_value.Value) Builder {
	app.val = val
	return app
}

// WithOperation adds an operation to the builder
func (app *builder) WithOperation(operation Operation) Builder {
	app.op = operation
	return app
}

// IsPrint flags the builder as print
func (app *builder) IsPrint() Builder {
	app.isPrint = true
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.val == nil {
		return nil, errors.New("the Value is mandatory in order to build a Value instruction instance")
	}

	if app.op != nil {
		return createValue(app.val, app.op), nil
	}

	builder := app.operationBuilder.Create()
	if app.isPrint {
		op, err := builder.IsPrint().Now()
		if err != nil {
			return nil, err
		}

		app.op = op
	}

	if app.op == nil {
		return nil, errors.New("the value is invalid because it does not contain an operation")
	}

	return createValue(app.val, app.op), nil
}
