package rules

import (
	"errors"
	"regexp"
)

type patternBuilder struct {
	code        string
	pattern     string
	possibility Possibility
}

func createPatternBuilder() PatternBuilder {
	out := patternBuilder{
		code:        "",
		pattern:     "",
		possibility: nil,
	}

	return &out
}

// Create initializes the builder
func (app *patternBuilder) Create() PatternBuilder {
	return createPatternBuilder()
}

// WithPattern adds a pattern to the builder
func (app *patternBuilder) WithPattern(pattern string) PatternBuilder {
	app.pattern = pattern
	return app
}

// WithCode adds a code to the builder
func (app *patternBuilder) WithCode(code string) PatternBuilder {
	app.code = code
	return app
}

// WithPossibility adds a possibility to the builder
func (app *patternBuilder) WithPossibility(possibility Possibility) PatternBuilder {
	app.possibility = possibility
	return app
}

// Now builds a new Pattern instance
func (app *patternBuilder) Now() (Pattern, error) {
	if app.pattern == "" {
		return nil, errors.New("the pattern is mandatory in order to build a Pattern instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Pattern instance")
	}

	if app.possibility == nil {
		return nil, errors.New("the possibility is mandatory in order to build a Pattern instance")
	}

	pattern, err := regexp.Compile(app.pattern)
	if err != nil {
		return nil, err
	}

	return createPattern(app.pattern, app.code, pattern, app.possibility), nil
}
