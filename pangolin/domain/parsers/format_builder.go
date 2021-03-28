package parsers

import "errors"

type formatBuilder struct {
	results string
	pattern string
	first   string
	second  string
}

func createFormatBuilder() FormatBuilder {
	out := formatBuilder{
		results: "",
		pattern: "",
		first:   "",
		second:  "",
	}

	return &out
}

// Create initializes the builder
func (app *formatBuilder) Create() FormatBuilder {
	return createFormatBuilder()
}

// WithResults adds a results variableName to the builder
func (app *formatBuilder) WithResults(results string) FormatBuilder {
	app.results = results
	return app
}

// WithPattern adds a pattern identifier to the builder
func (app *formatBuilder) WithPattern(pattern string) FormatBuilder {
	app.pattern = pattern
	return app
}

// WithFirst adds a first identifier to the builder
func (app *formatBuilder) WithFirst(first string) FormatBuilder {
	app.first = first
	return app
}

// WithSecond adds a second identifier to the builder
func (app *formatBuilder) WithSecond(second string) FormatBuilder {
	app.second = second
	return app
}

// Now builds a new Format instance
func (app *formatBuilder) Now() (Format, error) {
	if app.results == "" {
		return nil, errors.New("the results string is mandatory in order to build a Format instance")
	}

	if app.pattern == "" {
		return nil, errors.New("the pattern string is mandatory in order to build a Format instance")
	}

	if app.first == "" {
		return nil, errors.New("the first string is mandatory in order to build a Format instance")
	}

	if app.second == "" {
		return nil, errors.New("the second string is mandatory in order to build a Format instance")
	}

	return createFormat(app.results, app.pattern, app.first, app.second), nil
}
