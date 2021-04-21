package languages

import "errors"

type builder struct {
	variable string
	values   []Value
}

func createBuilder() Builder {
	out := builder{
		variable: "",
		values:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithValues add values to the builder
func (app *builder) WithValues(values []Value) Builder {
	app.values = values
	return app
}

// Now builds a new Language instance
func (app *builder) Now() (Language, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Language instance")
	}

	if app.values != nil && len(app.values) <= 0 {
		app.values = nil
	}

	if app.values == nil {
		return nil, errors.New("there must be at least 1 Value in order to build a Language instance")
	}

	return createLanguage(app.variable, app.values), nil
}
