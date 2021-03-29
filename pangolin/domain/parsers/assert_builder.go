package parsers

import "errors"

type assertBuilder struct {
	index     int
	condition string
}

func createAssertBuilder() AssertBuilder {
	out := assertBuilder{
		index:     -1,
		condition: "",
	}

	return &out
}

// Create initializes the builder
func (app *assertBuilder) Create() AssertBuilder {
	return createAssertBuilder()
}

// WithIndex adds an index to the builder
func (app *assertBuilder) WithIndex(index int) AssertBuilder {
	app.index = index
	return app
}

// WithCondition adds a condition to the builder
func (app *assertBuilder) WithCondition(condition string) AssertBuilder {
	app.condition = condition
	return app
}

// Now builds a new Assert instance
func (app *assertBuilder) Now() (Assert, error) {
	if app.index < 0 {
		return nil, errors.New("the assert is mandatory in order to build an Assert instance")
	}

	if app.condition != "" {
		return createAssertWithCondition(app.index, app.condition), nil
	}

	return createAssert(app.index), nil
}
