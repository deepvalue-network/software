package parsers

import "errors"

type indexBuilder struct {
	variable string
}

func createIndexBuilder() IndexBuilder {
	out := indexBuilder{
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *indexBuilder) Create() IndexBuilder {
	return createIndexBuilder()
}

// WithVariable adds a variable to the builder
func (app *indexBuilder) WithVariable(variable string) IndexBuilder {
	app.variable = variable
	return app
}

// Now builds a new Index instance
func (app *indexBuilder) Now() (Index, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Index instance")
	}

	return createIndex(app.variable), nil
}
