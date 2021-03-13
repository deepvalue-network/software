package parsers

import "errors"

type formatBuilder struct {
	results VariableName
	pattern Identifier
	first   Identifier
	second  Identifier
}

func createFormatBuilder() FormatBuilder {
	out := formatBuilder{
		results: nil,
		pattern: nil,
		first:   nil,
		second:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *formatBuilder) Create() FormatBuilder {
	return createFormatBuilder()
}

// WithResults adds a results variableName to the builder
func (app *formatBuilder) WithResults(results VariableName) FormatBuilder {
	app.results = results
	return app
}

// WithPattern adds a pattern identifier to the builder
func (app *formatBuilder) WithPattern(pattern Identifier) FormatBuilder {
	app.pattern = pattern
	return app
}

// WithFirst adds a first identifier to the builder
func (app *formatBuilder) WithFirst(first Identifier) FormatBuilder {
	app.first = first
	return app
}

// WithSecond adds a second identifier to the builder
func (app *formatBuilder) WithSecond(second Identifier) FormatBuilder {
	app.second = second
	return app
}

// Now builds a new Format instance
func (app *formatBuilder) Now() (Format, error) {
	if app.results == nil {
		return nil, errors.New("the results VariableName is mandatory in order to build a Format instance")
	}

	if app.pattern == nil {
		return nil, errors.New("the pattern Identifier is mandatory in order to build a Format instance")
	}

	if app.first == nil {
		return nil, errors.New("the first Identifier is mandatory in order to build a Format instance")
	}

	if app.second == nil {
		return nil, errors.New("the second Identifier is mandatory in order to build a Format instance")
	}

	return createFormat(app.results, app.pattern, app.first, app.second), nil
}
