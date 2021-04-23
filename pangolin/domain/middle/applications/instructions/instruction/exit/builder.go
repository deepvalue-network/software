package exit

type builder struct {
	condition string
}

func createBuilder() Builder {
	out := builder{
		condition: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition string) Builder {
	app.condition = condition
	return app
}

// Now builds a new Exit instance
func (app *builder) Now() (Exit, error) {
	if app.condition != "" {
		return createExitWithCondition(app.condition), nil
	}

	return createExit(), nil
}
