package geometries

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/shaders"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

type builder struct {
	id         *uuid.UUID
	isTriangle bool
	primitive  primitives.Primitive
	shader     shaders.Shader
	vertices   []vertices.Vertex
}

func createBuilder() Builder {
	out := builder{
		id:         nil,
		isTriangle: false,
		primitive:  nil,
		shader:     nil,
		vertices:   nil,
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

// WithPrimitive adds a primitive to the builder
func (app *builder) WithPrimitive(primitive primitives.Primitive) Builder {
	app.primitive = primitive
	return app
}

// WithVertices add vertices to the builder
func (app *builder) WithVertices(vertices []vertices.Vertex) Builder {
	app.vertices = vertices
	return app
}

// WithShader adds a shader to the builder
func (app *builder) WithShader(shader shaders.Shader) Builder {
	app.shader = shader
	return app
}

// IsTriangle flags the builder as triangles
func (app *builder) IsTriangle() Builder {
	app.isTriangle = true
	return app
}

// Now builds a new Geometry instance
func (app *builder) Now() (Geometry, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Geometry instance")
	}

	if app.primitive != nil {
		return nil, errors.New("finish the primitives loading in geometries")
	}

	if app.shader == nil {
		return nil, errors.New("the shader is mandatory in order to build a Geometry instance")
	}

	if app.vertices != nil && len(app.vertices) <= 0 {
		app.vertices = nil
	}

	if app.vertices == nil {
		return nil, errors.New("the vertices are mandatory in order to build a Geometry instance")
	}

	var typ Type
	if app.isTriangle {
		typ = createTypeWithTriangle()
	}

	if typ == nil {
		return nil, errors.New("the type is mandatory in order to build a Geometry instance")
	}

	return createGeometry(app.id, typ, app.shader, app.vertices), nil
}
