package tokens

import "errors"

type lineBuilder struct {
	elements []Element
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements []Element) LineBuilder {
	app.elements = elements
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.elements != nil && len(app.elements) <= 0 {
		app.elements = nil
	}

	if app.elements == nil {
		return nil, errors.New("the elements are mandatory in order to build a Line instance")
	}

	return createLine(app.elements), nil
}
