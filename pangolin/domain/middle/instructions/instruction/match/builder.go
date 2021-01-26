package match

import "errors"

type builder struct {
	input   string
	pattern string
}

func createBuilder() Builder {
	out := builder{
		input:   "",
		pattern: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInput adds the input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// WithPattern adds the pattern to the builder
func (app *builder) WithPattern(pattern string) Builder {
	app.pattern = pattern
	return app
}

// Now builds a new Match instance
func (app *builder) Now() (Match, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Match instance")
	}

	if app.pattern != "" {
		return createMatchWithPattern(app.input, app.pattern), nil
	}

	return createMatch(app.input), nil
}
