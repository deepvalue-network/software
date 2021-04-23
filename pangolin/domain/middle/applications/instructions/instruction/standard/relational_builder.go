package standard

import "errors"

type relationalBuilder struct {
	isLessThan bool
	isEqual    bool
	isNotEqual bool
}

func createRelationalBuilder() RelationalBuilder {
	out := relationalBuilder{
		isLessThan: false,
		isEqual:    false,
		isNotEqual: false,
	}

	return &out
}

// Create initializes the builder
func (app *relationalBuilder) Create() RelationalBuilder {
	return createRelationalBuilder()
}

// IsLessThan flags the builder as isLessThan
func (app *relationalBuilder) IsLessThan() RelationalBuilder {
	app.isLessThan = true
	return app
}

// IsEqual flags the builder as equal
func (app *relationalBuilder) IsEqual() RelationalBuilder {
	app.isEqual = true
	return app
}

// IsNotEqual flags the builder as notEqual
func (app *relationalBuilder) IsNotEqual() RelationalBuilder {
	app.isNotEqual = true
	return app
}

// Now builds a new Relational instance
func (app *relationalBuilder) Now() (Relational, error) {
	if app.isLessThan {
		return createRelationalWithLessThan(), nil
	}

	if app.isEqual {
		return createRelationalWithEqual(), nil
	}

	if app.isNotEqual {
		return createRelationalWithNotEqual(), nil
	}

	return nil, errors.New("the Relational is invalid")
}
