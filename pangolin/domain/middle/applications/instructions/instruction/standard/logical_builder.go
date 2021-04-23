package standard

import "errors"

type logicalBuilder struct {
	isAnd bool
	isOr  bool
}

func createLogicalBuilder() LogicalBuilder {
	out := logicalBuilder{
		isAnd: false,
		isOr:  false,
	}

	return &out
}

// Create initializes the builder
func (app *logicalBuilder) Create() LogicalBuilder {
	return createLogicalBuilder()
}

// IsAnd flags the builder as a and
func (app *logicalBuilder) IsAnd() LogicalBuilder {
	app.isAnd = true
	return app
}

// IsOr flags the builder as a or
func (app *logicalBuilder) IsOr() LogicalBuilder {
	app.isOr = true
	return app
}

// Now builds a new Logical instance
func (app *logicalBuilder) Now() (Logical, error) {
	if app.isAnd {
		return createLogicalWithAnd(), nil
	}

	if app.isOr {
		return createLogicalWithOr(), nil
	}

	return nil, errors.New("the Logical is invalid")
}
