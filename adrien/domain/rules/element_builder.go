package rules

import "errors"

type elementBuilder struct {
	content Content
	code    string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		content: nil,
		code:    "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithContent adds a content to the builder
func (app *elementBuilder) WithContent(content Content) ElementBuilder {
	app.content = content
	return app
}

// WithCode adds a code to the builder
func (app *elementBuilder) WithCode(code string) ElementBuilder {
	app.code = code
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Element instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build an Element instance")
	}

	return createElement(app.content, app.code), nil
}
