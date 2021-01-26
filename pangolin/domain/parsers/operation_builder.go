package parsers

import "errors"

type operationBuilder struct {
	arythmetic Arythmetic
	relational Relational
	logical    Logical
}

func createOperationalBuilder() OperationBuilder {
	out := operationBuilder{
		arythmetic: nil,
		relational: nil,
		logical:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationalBuilder()
}

// WithArythmetic adds an arythmetic operator to the builder
func (app *operationBuilder) WithArythmetic(arythmetic Arythmetic) OperationBuilder {
	app.arythmetic = arythmetic
	return app
}

// WithRelational adds a relational operator to the builder
func (app *operationBuilder) WithRelational(relational Relational) OperationBuilder {
	app.relational = relational
	return app
}

// WithLogical adds a logical operator to the builder
func (app *operationBuilder) WithLogical(logical Logical) OperationBuilder {
	app.logical = logical
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.arythmetic != nil {
		return createOperationWithArythmetic(app.arythmetic), nil
	}

	if app.relational != nil {
		return createOperationWithRelational(app.relational), nil
	}

	if app.logical != nil {
		return createOperationWithLogical(app.logical), nil
	}

	return nil, errors.New("the Operation is invalid")
}
