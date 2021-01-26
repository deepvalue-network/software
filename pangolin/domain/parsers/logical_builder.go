package parsers

import "errors"

type logicalBuilder struct {
	and StandardOperation
	or  StandardOperation
}

func createLogicalBuilder() LogicalBuilder {
	out := logicalBuilder{
		and: nil,
		or:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *logicalBuilder) Create() LogicalBuilder {
	return createLogicalBuilder()
}

// WithAnd adds the and logical operator to the builder
func (app *logicalBuilder) WithAnd(and StandardOperation) LogicalBuilder {
	app.and = and
	return app
}

// WithOr adds the or logical operator to the builder
func (app *logicalBuilder) WithOr(or StandardOperation) LogicalBuilder {
	app.or = or
	return app
}

// Now builds a new Logical instance
func (app *logicalBuilder) Now() (Logical, error) {
	if app.and != nil {
		return createLogicalWithAnd(app.and), nil
	}

	if app.or != nil {
		return createLogicalWithOr(app.or), nil
	}

	return nil, errors.New("the Logical instance is invalid")
}
