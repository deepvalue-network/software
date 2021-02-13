package layers

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/alphas"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures"
)

type builder struct {
	id    *uuid.UUID
	index uint
	alpha alphas.Alpha
	tex   textures.Texture
}

func createBuilder() Builder {
	out := builder{
		id:    nil,
		index: 0,
		alpha: nil,
		tex:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
	return app
}

// WithAlpha adds an alpha to the builder
func (app *builder) WithAlpha(alpha alphas.Alpha) Builder {
	app.alpha = alpha
	return app
}

// WithTexture adds a texture to the builder
func (app *builder) WithTexture(tex textures.Texture) Builder {
	app.tex = tex
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Layer instance")
	}

	if app.alpha == nil {
		return nil, errors.New("the alpha is mandatory in order to build a Layer instance")
	}

	if app.tex == nil {
		return nil, errors.New("the texture is mandatory in order to build a Layer instance")
	}

	return createLayer(app.id, app.index, app.alpha, app.tex), nil
}
