package value

import "errors"

type operationBuilder struct {
	isPrint bool
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		isPrint: false,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// IsPrint flags the operation as print
func (app *operationBuilder) IsPrint() OperationBuilder {
	app.isPrint = true
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.isPrint {
		return createOperationWithPrint(), nil
	}

	return nil, errors.New("the Operation is invalid")
}
