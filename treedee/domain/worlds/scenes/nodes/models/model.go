package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

type model struct {
	id  *uuid.UUID
	geo geometries.Geometry
	mat materials.Material
}

func createModel(
	id *uuid.UUID,
	geo geometries.Geometry,
	mat materials.Material,
) Model {
	out := model{
		id:  id,
		geo: geo,
		mat: mat,
	}

	return &out
}

// ID returns the id
func (obj *model) ID() *uuid.UUID {
	return obj.id
}

// Geometry returns the geometry
func (obj *model) Geometry() geometries.Geometry {
	return obj.geo
}

// Material returns the material
func (obj *model) Material() materials.Material {
	return obj.mat
}
