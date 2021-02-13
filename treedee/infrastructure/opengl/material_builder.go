package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

type materialBuilder struct {
	layerBuilder LayerBuilder
	mat          materials.Material
}

func createMaterialBuilder(
	layerBuilder LayerBuilder,
) MaterialBuilder {
	out := materialBuilder{
		layerBuilder: layerBuilder,
		mat:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *materialBuilder) Create() MaterialBuilder {
	return createMaterialBuilder(app.layerBuilder)
}

// WithMaterial adds a material to the builder
func (app *materialBuilder) WithMaterial(mat materials.Material) MaterialBuilder {
	app.mat = mat
	return app
}

// Now builds a new Material instance
func (app *materialBuilder) Now() (Material, error) {
	if app.mat == nil {
		return nil, errors.New("the material is mandatory in order to build a Material instance")
	}

	layers := []Layer{}
	domainLayers := app.mat.Layers()
	for _, oneDomainLayer := range domainLayers {
		layer, err := app.layerBuilder.Create().WithLayer(oneDomainLayer).Now()
		if err != nil {
			return nil, err
		}

		layers = append(layers, layer)
	}

	domainAlpha := app.mat.Alpha()
	value := domainAlpha.Alpha()
	normalizedValue := float32(value / maxAlpha)
	variable := domainAlpha.Variable()
	alpha := createAlpha(normalizedValue, variable)

	id := app.mat.ID()
	return createMaterial(id, alpha, app.mat.Viewport(), layers), nil
}
