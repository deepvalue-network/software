package opengl

import uuid "github.com/satori/go.uuid"

type vertexShader struct {
	id        *uuid.UUID
	variables VertexShaderVariables
}

func createVertexShader(
	id *uuid.UUID,
	variables VertexShaderVariables,
) VertexShader {
	out := vertexShader{
		id:        id,
		variables: variables,
	}

	return &out
}

// ID returns the id
func (obj *vertexShader) ID() *uuid.UUID {
	return obj.id
}

// Variables returns the variables
func (obj *vertexShader) Variables() VertexShaderVariables {
	return obj.variables
}
