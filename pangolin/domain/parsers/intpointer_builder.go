package parsers

import "errors"

type intPointerBuilder struct {
	intVal   int64
	variable string
}

func createIntPointerBuilder() IntPointerBuilder {
	out := intPointerBuilder{
		intVal:   -1,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *intPointerBuilder) Create() IntPointerBuilder {
	return createIntPointerBuilder()
}

// WithInt adds an int to the builder
func (app *intPointerBuilder) WithInt(intVal int64) IntPointerBuilder {
	app.intVal = intVal
	return app
}

// WithVariable adds a variable to the builder
func (app *intPointerBuilder) WithVariable(variable string) IntPointerBuilder {
	app.variable = variable
	return app
}

// Now builds a new IntPointer instance
func (app *intPointerBuilder) Now() (IntPointer, error) {
	if app.intVal >= 0 {
		return createIntPointerWithInt(app.intVal), nil
	}

	if app.variable != "" {
		return createIntPointerWithVariable(app.variable), nil
	}

	return nil, errors.New("the IntPointer is invalid")
}
