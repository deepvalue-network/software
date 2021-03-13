package format

import "errors"

type builder struct {
	results string
	pattern string
	first   string
	second  string
}

func createBuilder() Builder {
	out := builder{
		results: "",
		pattern: "",
		first:   "",
		second:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResults add results to the builder
func (app *builder) WithResults(results string) Builder {
	app.results = results
	return app
}

// WithPattern add pattern to the builder
func (app *builder) WithPattern(pattern string) Builder {
	app.pattern = pattern
	return app
}

// WithFirst add first to the builder
func (app *builder) WithFirst(first string) Builder {
	app.first = first
	return app
}

// WithSecond add second to the builder
func (app *builder) WithSecond(second string) Builder {
	app.second = second
	return app
}

// Now builds a new Format instance
func (app *builder) Now() (Format, error) {
	if app.results == "" {
		return nil, errors.New("the results is mandatory in order to build a Format instance")
	}

	if app.pattern == "" {
		return nil, errors.New("the pattern is mandatory in order to build a Format instance")
	}

	if app.first == "" {
		return nil, errors.New("the first is mandatory in order to build a Format instance")
	}

	if app.second == "" {
		return nil, errors.New("the second is mandatory in order to build a Format instance")
	}

	return createFormat(app.results, app.pattern, app.first, app.second), nil
}
