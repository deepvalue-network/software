package middle

import "errors"

type patternMatchBuilder struct {
	pattern string
	enter   string
	exit    string
}

func createPatternMatchBuilder() PatternMatchBuilder {
	out := patternMatchBuilder{
		pattern: "",
		enter:   "",
		exit:    "",
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

// WithEnterLabel adds an enter label to the builder
func (app *patternMatchBuilder) WithEnterLabel(enter string) PatternMatchBuilder {
	app.enter = enter
	return app
}

// WithExitLabel adds an exit label to the builder
func (app *patternMatchBuilder) WithExitLabel(exit string) PatternMatchBuilder {
	app.exit = exit
	return app
}

// Now builds a new PatternMatch instance
func (app *patternMatchBuilder) Now() (PatternMatch, error) {
	if app.pattern == "" {
		return nil, errors.New("the pattern is mandatory in order to build a PatternMatch instance")
	}

	if app.enter != "" && app.exit != "" {
		return createPatternMatchWithEnterAndExit(app.pattern, app.enter, app.exit), nil
	}

	if app.enter != "" {
		return createPatternMatchWithEnter(app.pattern, app.enter), nil
	}

	if app.exit != "" {
		return createPatternMatchWithExit(app.pattern, app.exit), nil
	}

	return nil, errors.New("the PatternMatch is invalid")
}
