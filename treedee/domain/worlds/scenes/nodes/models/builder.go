package models

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

type builder struct {
	id  *uuid.UUID
	geo geometries.Geometry
	mat materials.Material
}

func createBuilder() Builder {
	out := builder{
		id:  nil,
		geo: nil,
		mat: nil,
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

// WithGeometry adds a geometry to the builder
func (app *builder) WithGeometry(geo geometries.Geometry) Builder {
	app.geo = geo
	return app
}

// WithMaterial adds a material to the builder
func (app *builder) WithMaterial(material materials.Material) Builder {
	app.mat = material
	return app
}

// Now builds a new Model instance
func (app *builder) Now() (Model, error) {
	if app.id == nil {
		return nil, errors.New("the id is mandatory in order to build a Model instance")
	}

	if app.geo == nil {
		return nil, errors.New("the geometry is mandatory in order to build a Model instance")
	}

	if app.mat == nil {
		return nil, errors.New("the material is mandatory in order to build a Model instance")
	}

	return createModel(app.id, app.geo, app.mat), nil
}
