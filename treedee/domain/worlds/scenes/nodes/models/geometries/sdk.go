package geometries

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the geometry builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithPrimitive(primitive primitives.Primitive) Builder
	WithVertices(vertices []vertices.Vertex) Builder
	WithShader(shader shaders.Shader) Builder
	IsTriangle() Builder
	Now() (Geometry, error)
}

// Geometry represents a geometry
type Geometry interface {
	ID() *uuid.UUID
	Type() Type
	Shader() shaders.Shader
	Vertices() []vertices.Vertex
}

// Type represents geometry type
type Type interface {
	IsTriangle() bool
}
