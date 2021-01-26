package parsers

import "errors"

type relationalBuilder struct {
	lessThan StandardOperation
	equal    StandardOperation
	notEqual StandardOperation
}

func createRelationalBuilder() RelationalBuilder {
	out := relationalBuilder{
		lessThan: nil,
		equal:    nil,
		notEqual: nil,
	}

	return &out
}

// Create initializes the builder
func (app *relationalBuilder) Create() RelationalBuilder {
	return createRelationalBuilder()
}

// WithLessThan adds a lessThan relational operator to the builder
func (app *relationalBuilder) WithLessThan(lessThan StandardOperation) RelationalBuilder {
	app.lessThan = lessThan
	return app
}

// WithEqual adds an equal relational operator to the builder
func (app *relationalBuilder) WithEqual(equal StandardOperation) RelationalBuilder {
	app.equal = equal
	return app
}

// WithNotEqual adds a notEqual relational operator to the builder
func (app *relationalBuilder) WithNotEqual(notEqual StandardOperation) RelationalBuilder {
	app.notEqual = notEqual
	return app
}

// Now builds a new Relational instance
func (app *relationalBuilder) Now() (Relational, error) {
	if app.lessThan != nil {
		return createRelationalWithLessThan(app.lessThan), nil
	}

	if app.equal != nil {
		return createRelationalWithEqual(app.equal), nil
	}

	if app.notEqual != nil {
		return createRelationalWithNotEqual(app.notEqual), nil
	}

	return nil, errors.New("the Relational instance is invalid")
}
