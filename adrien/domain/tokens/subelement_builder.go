package tokens

import "errors"

type subElementBuilder struct {
	content     Content
	cardinality SpecificCardinality
}

func createSubElementBuilder() SubElementBuilder {
	out := subElementBuilder{
		content:     nil,
		cardinality: nil,
	}

	return &out
}

// Create initializes the builder
func (app *subElementBuilder) Create() SubElementBuilder {
	return createSubElementBuilder()
}

// WithContent adds content to the builder
func (app *subElementBuilder) WithContent(content Content) SubElementBuilder {
	app.content = content
	return app
}

// WithCardinality adds cardinality to the builder
func (app *subElementBuilder) WithCardinality(cardinality SpecificCardinality) SubElementBuilder {
	app.cardinality = cardinality
	return app
}

// Now builds a new SubElement instance
func (app *subElementBuilder) Now() (SubElement, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a SubElement instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the specificCardinality is mandatory in order to build a SubElement instance")
	}

	return createSubElement(app.content, app.cardinality), nil
}
