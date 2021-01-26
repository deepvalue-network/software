package parsers

type exitBuilder struct {
	condition Identifier
}

func createExitBuilder() ExitBuilder {
	out := exitBuilder{
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *exitBuilder) Create() ExitBuilder {
	return createExitBuilder()
}

// WithCondition adds a condition to the builder
func (app *exitBuilder) WithCondition(cond Identifier) ExitBuilder {
	app.condition = cond
	return app
}

// Now builds a new Exit instance
func (app *exitBuilder) Now() (Exit, error) {
	if app.condition != nil {
		return createExitWithCondition(app.condition), nil
	}

	return createExit(), nil
}
