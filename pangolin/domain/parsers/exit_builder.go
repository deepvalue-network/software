package parsers

type exitBuilder struct {
	condition string
}

func createExitBuilder() ExitBuilder {
	out := exitBuilder{
		condition: "",
	}

	return &out
}

// Create initializes the builder
func (app *exitBuilder) Create() ExitBuilder {
	return createExitBuilder()
}

// WithCondition adds a condition to the builder
func (app *exitBuilder) WithCondition(cond string) ExitBuilder {
	app.condition = cond
	return app
}

// Now builds a new Exit instance
func (app *exitBuilder) Now() (Exit, error) {
	if app.condition != "" {
		return createExitWithCondition(app.condition), nil
	}

	return createExit(), nil
}
