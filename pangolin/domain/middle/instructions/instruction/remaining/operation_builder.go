package remaining

import "errors"

type operationBuilder struct {
	ary  Arythmetic
	misc Misc
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		ary:  nil,
		misc: nil,
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

// WithMisc adds a misc to the builder
func (app *operationBuilder) WithMisc(misc Misc) OperationBuilder {
	app.misc = misc
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.ary != nil {
		return createOperationWithArythmetic(app.ary), nil
	}

	if app.misc != nil {
		return createOperationWithMisc(app.misc), nil
	}

	return nil, errors.New("the Operation is invalid")
}
