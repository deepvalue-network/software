package geometries

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

type geometry struct {
	id       *uuid.UUID
	typ      Type
	shader   shaders.Shader
	vertices []vertices.Vertex
}

func createGeometry(
	id *uuid.UUID,
	typ Type,
	shader shaders.Shader,
	vertices []vertices.Vertex,
) Geometry {
	out := geometry{
		id:       id,
		typ:      typ,
		shader:   shader,
		vertices: vertices,
	}

	return &out
}

// ID returns the id
func (obj *geometry) ID() *uuid.UUID {
	return obj.id
}

// Type returns the type
func (obj *geometry) Type() Type {
	return obj.typ
}

// Shader returns the shader
func (obj *geometry) Shader() shaders.Shader {
	return obj.shader
}

// Vertices returns the vertices
func (obj *geometry) Vertices() []vertices.Vertex {
	return obj.vertices
}
