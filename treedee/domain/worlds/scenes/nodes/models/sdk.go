package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the model builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithGeometry(geo geometries.Geometry) Builder
	WithMaterial(material materials.Material) Builder
	Now() (Model, error)
}

// Model represents a model
type Model interface {
	ID() *uuid.UUID
	Geometry() geometries.Geometry
	Material() materials.Material
}
