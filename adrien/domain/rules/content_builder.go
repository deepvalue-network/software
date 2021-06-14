package rules

import "errors"

type contentBuilder struct {
	constant string
	pattern  Pattern
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		constant: "",
		pattern:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithConstant adds a constant to the builder
func (app *contentBuilder) WithConstant(constant string) ContentBuilder {
	app.constant = constant
	return app
}

// WithPattern adds a pattern to the builder
func (app *contentBuilder) WithPattern(pattern Pattern) ContentBuilder {
	app.pattern = pattern
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.constant != "" {
		return createContentWithConstant(app.constant), nil
	}

	if app.pattern != nil {
		return createContentWithPattern(app.pattern), nil
	}

	return nil, errors.New("the Content is invalid")
}
