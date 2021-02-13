package shaders

import uuid "github.com/satori/go.uuid"

// NewBuilder returns the builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a material shader
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithCode(code string) Builder
	Now() (Shader, error)
}

// Shader represents a material shader
type Shader interface {
	ID() *uuid.UUID
	Code() string
}
