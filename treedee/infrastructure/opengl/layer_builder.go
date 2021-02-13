package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers"
)

type layerBuilder struct {
	textureBuilder TextureBuilder
	layer          layers.Layer
}

func createLayerBuilder(
	textureBuilder TextureBuilder,
) LayerBuilder {
	out := layerBuilder{
		textureBuilder: textureBuilder,
		layer:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerBuilder) Create() LayerBuilder {
	return createLayerBuilder(
		app.textureBuilder,
	)
}

// WithLayer adds a layer to the builder
func (app *layerBuilder) WithLayer(layer layers.Layer) LayerBuilder {
	app.layer = layer
	return app
}

// Now builds a new Layer instance
func (app *layerBuilder) Now() (Layer, error) {
	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Layer instance")
	}

	domainTex := app.layer.Texture()
	tex, err := app.textureBuilder.Create().WithTexture(domainTex).Now()
	if err != nil {
		return nil, err
	}

	domainAlpha := app.layer.Alpha()
	value := domainAlpha.Alpha()
	normalizedValue := float32(value / maxAlpha)
	variable := domainAlpha.Variable()
	alpha := createAlpha(normalizedValue, variable)

	id := app.layer.ID()
	index := app.layer.Index()
	return createLayer(id, index, alpha, tex), nil
}
