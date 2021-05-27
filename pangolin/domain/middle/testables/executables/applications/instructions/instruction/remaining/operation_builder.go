package remaining

import "errors"

type operationBuilder struct {
	ary Arythmetic
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		ary: nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// WithArythmetic adds an arythmetic to the builder
func (app *operationBuilder) WithArythmetic(ary Arythmetic) OperationBuilder {
	app.ary = ary
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.ary != nil {
		return createOperationWithArythmetic(app.ary), nil
	}

	return nil, errors.New("the Operation is invalid")
}
