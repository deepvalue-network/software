package registry

import "errors"

type indexBuilder struct {
	intVal   int64
	variable string
}

func createIndexBuilder() IndexBuilder {
	out := indexBuilder{
		intVal:   -1,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *indexBuilder) Create() IndexBuilder {
	return createIndexBuilder()
}

// WithInt adds an int to the builder
func (app *indexBuilder) WithInt(intVal int64) IndexBuilder {
	app.intVal = intVal
	return app
}

// WithVariable adds a variable to the builder
func (app *indexBuilder) WithVariable(variable string) IndexBuilder {
	app.variable = variable
	return app
}

// Now builds a new Index instance
func (app *indexBuilder) Now() (Index, error) {
	if app.intVal >= 0 {
		return createIndexWithInt(app.intVal), nil
	}

	if app.variable != "" {
		return createIndexWithVariable(app.variable), nil
	}

	return nil, errors.New("the Index instance is invalid")
}
