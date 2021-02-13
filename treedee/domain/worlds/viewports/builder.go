package viewports

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
)

type builder struct {
	rectangle ints.Rectangle
	variable  string
}

func createBuilder() Builder {
	out := builder{
		rectangle: nil,
		variable:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRectangle adds a rectangle to the builder
func (app *builder) WithRectangle(rect ints.Rectangle) Builder {
	app.rectangle = rect
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// Now builds a new Viewport instance
func (app *builder) Now() (Viewport, error) {
	if app.rectangle == nil {
		return nil, errors.New("the rectangle is mandatory in order to build a Viewport instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Viewport instance")
	}

	return createViewport(app.rectangle, app.variable), nil
}
