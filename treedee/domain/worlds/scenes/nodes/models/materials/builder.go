package materials

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/treedee/domain/worlds/alphas"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
	uuid "github.com/satori/go.uuid"
)

type builder struct {
	id       *uuid.UUID
	alpha    alphas.Alpha
	shader   shaders.Shader
	viewport viewports.Viewport
	layers   []layers.Layer
}

func createBuilder() Builder {
	out := builder{
		id:       nil,
		alpha:    nil,
		shader:   nil,
		viewport: nil,
		layers:   nil,
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

// WithAlpha adds an alpha to the builder
func (app *builder) WithAlpha(alpha alphas.Alpha) Builder {
	app.alpha = alpha
	return app
}

// WithShader adds a shader to the builder
func (app *builder) WithShader(shader shaders.Shader) Builder {
	app.shader = shader
	return app
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport viewports.Viewport) Builder {
	app.viewport = viewport
	return app
}

// WithLayers add layers to the builder
func (app *builder) WithLayers(layers []layers.Layer) Builder {
	app.layers = layers
	return app
}

// Now builds a new Material instance
func (app *builder) Now() (Material, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Material instance")
	}

	if app.alpha == nil {
		return nil, errors.New("the alpha is mandatory in order to build a Material instance")
	}

	if app.shader == nil {
		return nil, errors.New("the shader is mandatory in order to build a Material instance")
	}

	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Material instance")
	}

	if app.layers != nil && len(app.layers) <= 0 {
		app.layers = nil
	}

	if app.layers == nil {
		return nil, errors.New("there must be at least 1 Layer in order to build a Material instance")
	}

	var dim *ints.Vec2
	reOrdered := []layers.Layer{}
	for index, oneLayer := range app.layers {
		if dim == nil {
			texDim := oneLayer.Texture().Dimension()
			dim = &texDim
		}

		// make sure all layers have the same dimensions:
		layerDim := oneLayer.Texture().Dimension()
		if !layerDim.Compare(*dim) {
			str := fmt.Sprintf("there layer (index: %d) was expected to be of dimension %s, %s provided", index, dim.String(), layerDim.String())
			return nil, errors.New(str)
		}

		reOrdered[oneLayer.Index()] = oneLayer
	}

	// make sure the viewport can be contained within the textures:
	if app.viewport.IsContained(*dim) {
		str := fmt.Sprintf("the given viewport (%s) cannot be contained in the material's original dimensions: %s", app.viewport.Rectangle().String(), dim.String())
		return nil, errors.New(str)
	}

	if len(reOrdered) != len(app.layers) {
		str := fmt.Sprintf("the re-ordered layers (amount: %d) do not match the amount of original layers (%d), the indexes may be overlapping in the layers list", len(reOrdered), len(app.layers))
		return nil, errors.New(str)
	}

	return createMaterial(app.id, app.alpha, app.shader, app.viewport, reOrdered), nil
}
