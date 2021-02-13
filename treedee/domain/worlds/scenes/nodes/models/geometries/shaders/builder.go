package shaders

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type builder struct {
	id     *uuid.UUID
	code   string
	tex    string
	vertex string
}

func createBuilder() Builder {
	out := builder{
		id:     nil,
		code:   "",
		tex:    "",
		vertex: "",
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

// WithCode adds a code to the builder
func (app *builder) WithCode(code string) Builder {
	app.code = code
	return app
}

// WithTextureCoordinatesVariable adds a texture coordinates variable to the builder
func (app *builder) WithTextureCoordinatesVariable(texCoordVar string) Builder {
	app.tex = texCoordVar
	return app
}

// WithVertexCoordinatesVariable adds a vertex coordinates variable to the builder
func (app *builder) WithVertexCoordinatesVariable(verCoordVar string) Builder {
	app.vertex = verCoordVar
	return app
}

// Now builds a new Shader instance
func (app *builder) Now() (Shader, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Shader instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Shader instance")
	}

	if app.tex == "" {
		return nil, errors.New("the texture coordinates variable is mandatory in order to build a Shader instance")
	}

	if app.vertex == "" {
		return nil, errors.New("the vertex coordinates variable is mandatory in order to build a Shader instance")
	}

	variables := createVariables(app.tex, app.vertex)
	return createShader(app.id, app.code, variables), nil
}
