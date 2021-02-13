package ints

import "errors"

type builder struct {
	pos *Vec2
	dim *Vec2
}

func createBuilder() Builder {
	out := builder{
		pos: nil,
		dim: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(pos Vec2) Builder {
	app.pos = &pos
	return app
}

// WithDimension adds a dimension to the builder
func (app *builder) WithDimension(dim Vec2) Builder {
	app.dim = &dim
	return app
}

// Now builds a new Rectangle instance
func (app *builder) Now() (Rectangle, error) {
	if app.pos == nil {
		return nil, errors.New("the position is mandatory in order to build a Rectangle instance")
	}

	if app.dim == nil {
		return nil, errors.New("the dimension is mandatory in order to build a Rectangle instance")
	}

	return createRectangle(*app.pos, *app.dim), nil
}
