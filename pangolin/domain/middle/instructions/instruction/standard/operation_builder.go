package standard

import "errors"

type operationBuilder struct {
	ary  Arythmetic
	rel  Relational
	log  Logical
	misc Misc
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		ary:  nil,
		rel:  nil,
		log:  nil,
		misc: nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// WithArythmetic adds an arythmetic to the builder
func (app *operationBuilder) WithArythmetic(arythmetic Arythmetic) OperationBuilder {
	app.ary = arythmetic
	return app
}

// WithRelational adds a relational to the builder
func (app *operationBuilder) WithRelational(relational Relational) OperationBuilder {
	app.rel = relational
	return app
}

// WithLogical adds a logical to the builder
func (app *operationBuilder) WithLogical(logical Logical) OperationBuilder {
	app.log = logical
	return app
}

// WithMisc adds a misc to the builder
func (app *operationBuilder) WithMisc(misc Misc) OperationBuilder {
	app.misc = misc
	return app
}

// Now builds a new Operaton instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.ary != nil {
		return createOperationWithArythmetic(app.ary), nil
	}

	if app.rel != nil {
		return createOperationWithRelational(app.rel), nil
	}

	if app.log != nil {
		return createOperationWithLogical(app.log), nil
	}

	if app.misc != nil {
		return createOperationWithMisc(app.misc), nil
	}

	return nil, errors.New("the Operation is invalid")
}
