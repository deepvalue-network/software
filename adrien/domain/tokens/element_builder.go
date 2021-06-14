package tokens

import "errors"

type elementBuilder struct {
	content     Content
	code        string
	subElements SubElements
	cardinality Cardinality
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		content:     nil,
		code:        "",
		subElements: nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithContent adds content to the builder
func (app *elementBuilder) WithContent(content Content) ElementBuilder {
	app.content = content
	return app
}

// WithCode adds code to the builder
func (app *elementBuilder) WithCode(code string) ElementBuilder {
	app.code = code
	return app
}

// WithSubElements add subElements to the builder
func (app *elementBuilder) WithSubElements(subElements SubElements) ElementBuilder {
	app.subElements = subElements
	return app
}

// WithCardinality add cardinality to the builder
func (app *elementBuilder) WithCardinality(cardinality Cardinality) ElementBuilder {
	app.cardinality = cardinality
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

	if app.subElements != nil && app.cardinality != nil {
		return createElementWithSubElementsAndCardinality(app.content, app.code, app.subElements, app.cardinality), nil
	}

	if app.subElements != nil {
		return createElementWithSubElements(app.content, app.code, app.subElements), nil
	}

	if app.cardinality != nil {
		return createElementWithCardinality(app.content, app.code, app.cardinality), nil
	}

	return createElement(app.content, app.code), nil
}
