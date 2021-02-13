package shaders

import uuid "github.com/satori/go.uuid"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a shader builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithCode(code string) Builder
	WithTextureCoordinatesVariable(texCoordVar string) Builder
	WithVertexCoordinatesVariable(verCoordVar string) Builder
	Now() (Shader, error)
}

// Shader represents a vertex shader
type Shader interface {
	ID() *uuid.UUID
	Code() string
	Variables() Variables
}

// Variables represents shader variables
type Variables interface {
	TextureCoordinates() string
	VertexCoordinates() string
}
