package parsers

import "errors"

type patternMatchBuilder struct {
	pattern  string
	labels   PatternLabels
	variable string
}

func createPatternMatchBuilder() PatternMatchBuilder {
	out := patternMatchBuilder{
		pattern:  "",
		labels:   nil,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *patternMatchBuilder) Create() PatternMatchBuilder {
	return createPatternMatchBuilder()
}

// WithPattern adds a pattern to the builder
func (app *patternMatchBuilder) WithPattern(pattern string) PatternMatchBuilder {
	app.pattern = pattern
	return app
}

// WithLabels adds patternLabels to the builder
func (app *patternMatchBuilder) WithLabels(labels PatternLabels) PatternMatchBuilder {
	app.labels = labels
	return app
}

// WithVariable adds variable name to the builder
func (app *patternMatchBuilder) WithVariable(variable string) PatternMatchBuilder {
	app.variable = variable
	return app
}

// Now builds a new PatternMatch instance
func (app *patternMatchBuilder) Now() (PatternMatch, error) {
	if app.pattern == "" {
		return nil, errors.New("the pattern is mandatory in order to build a PatternMatch instance")
	}

	if app.labels == nil {
		return nil, errors.New("the patternLabels is mandatory in order to build a PatternMatch instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a PatternMatch instance")
	}

	return createPatternMatch(app.pattern, app.labels, app.variable), nil
}
