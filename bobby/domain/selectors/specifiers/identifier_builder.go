package specifiers

import "errors"

type identifierBuilder struct {
	element  Element
	comparer Comparer
}

func createIdentifierBuilder() IdentifierBuilder {
	out := identifierBuilder{
		element:  nil,
		comparer: nil,
	}

	return &out
}

// Create initializes the builder
func (app *identifierBuilder) Create() IdentifierBuilder {
	return createIdentifierBuilder()
}

// WithElement adds an element to the builder
func (app *identifierBuilder) WithElement(element Element) IdentifierBuilder {
	app.element = element
	return app
}

// WithComparer adds a comparer to the builder
func (app *identifierBuilder) WithComparer(comparer Comparer) IdentifierBuilder {
	app.comparer = comparer
	return app
}

// Now builds a new Identifier instance
func (app *identifierBuilder) Now() (Identifier, error) {
	if app.element != nil {
		return createIdentifierWithElement(app.element), nil
	}

	if app.comparer != nil {
		return createIdentifierWithComparer(app.comparer), nil
	}

	return nil, errors.New("the Identifier is invalid")
}
