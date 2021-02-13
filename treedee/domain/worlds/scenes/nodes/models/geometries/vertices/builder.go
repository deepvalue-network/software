package vertices

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
)

type builder struct {
	space *fl32.Vec3
	tex   *fl32.Vec2
}

func createBuilder() Builder {
	out := builder{
		space: nil,
		tex:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSpace adds a space to the builder
func (app *builder) WithSpace(space fl32.Vec3) Builder {
	app.space = &space
	return app
}

// WithTexture adds a texture to the builder
func (app *builder) WithTexture(tex fl32.Vec2) Builder {
	app.tex = &tex
	return app
}

// Now builds a new Vertex instance
func (app *builder) Now() (Vertex, error) {
	if app.space == nil {
		return nil, errors.New("the space is mandatory in order to build a Vertex instance")
	}

	if app.tex == nil {
		return nil, errors.New("the texture coordinates are mandatory in order to build a Vertex instance")
	}

	return createVertex(*app.space, *app.tex), nil
}
