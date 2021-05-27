package stackframe

import "errors"

type skipBuilder struct {
	intVal   int64
	variable string
}

func createSkipBuilder() SkipBuilder {
	out := skipBuilder{
		intVal:   -1,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *skipBuilder) Create() SkipBuilder {
	return createSkipBuilder()
}

// WithInt adds an int to the builder
func (app *skipBuilder) WithInt(intVal int64) SkipBuilder {
	app.intVal = intVal
	return app
}

// WithVariable adds a variable to the builder
func (app *skipBuilder) WithVariable(variable string) SkipBuilder {
	app.variable = variable
	return app
}

// Now builds a new Skip instance
func (app *skipBuilder) Now() (Skip, error) {
	if app.intVal >= 0 {
		return createSkipWithInt(app.intVal), nil
	}

	if app.variable != "" {
		return createSkipWithVariable(app.variable), nil
	}

	return nil, errors.New("the Skip instance is invalid")
}
