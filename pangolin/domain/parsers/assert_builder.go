package parsers

type assertBuilder struct {
	condition Identifier
}

func createAssertBuilder() AssertBuilder {
	out := assertBuilder{
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assertBuilder) Create() AssertBuilder {
	return createAssertBuilder()
}

// WithCondition adds a condition to the builder
func (app *assertBuilder) WithCondition(condition Identifier) AssertBuilder {
	app.condition = condition
	return app
}

// Now builds a new Assert instance
func (app *assertBuilder) Now() (Assert, error) {
	if app.condition != nil {
		return createAssertWithCondition(app.condition), nil
	}

	return createAssert(), nil
}
