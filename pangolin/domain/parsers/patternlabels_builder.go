package parsers

import "errors"

type patternLabelsBuilder struct {
	enter string
	exit  string
}

func createPatternLabelsBuilder() PatternLabelsBuilder {
	out := patternLabelsBuilder{
		enter: "",
		exit:  "",
	}

	return &out
}

// Create initializes the builder
func (app *patternLabelsBuilder) Create() PatternLabelsBuilder {
	return createPatternLabelsBuilder()
}

// WithEnterLabel adds an enter label to the builder
func (app *patternLabelsBuilder) WithEnterLabel(enter string) PatternLabelsBuilder {
	app.enter = enter
	return app
}

// WithExitLabel adds an exit label to the builder
func (app *patternLabelsBuilder) WithExitLabel(exit string) PatternLabelsBuilder {
	app.exit = exit
	return app
}

// Now builds a new PatternLabels instance
func (app *patternLabelsBuilder) Now() (PatternLabels, error) {
	if app.enter != "" && app.exit != "" {
		return createPatternLabelsWithEnterLabelAndExitLabel(app.enter, app.exit), nil
	}

	if app.enter != "" {
		return createPatternLabelsWithEnterLabel(app.enter), nil
	}

	if app.exit != "" {
		return createPatternLabelsWithExitLabel(app.exit), nil
	}

	return nil, errors.New("the enter or exit label is mandatory in order to build a PatternLabels instance")
}
