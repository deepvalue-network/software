package parsers

import "errors"

type matchBuilder struct {
	input   string
	pattern string
}

func createMatchBuilder() MatchBuilder {
	out := matchBuilder{
		input:   "",
		pattern: "",
	}

	return &out
}

// Create initializes the builder
func (app *matchBuilder) Create() MatchBuilder {
	return createMatchBuilder()
}

// WithInput adds an input to the builder
func (app *matchBuilder) WithInput(input string) MatchBuilder {
	app.input = input
	return app
}

// WithPattern adds a pattern to the builder
func (app *matchBuilder) WithPattern(pattern string) MatchBuilder {
	app.pattern = pattern
	return app
}

// Now builds a new Match instance
func (app *matchBuilder) Now() (Match, error) {
	if app.input == "" {
		return nil, errors.New("the input string is mandatory in order to build a Match instance")
	}

	if app.pattern != "" {
		return createMatchWithPattern(app.input, app.pattern), nil
	}

	return createMatch(app.input), nil
}
