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
	IsDynamic() Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	ID() *uuid.UUID
	Code() string
	IsDynamic() bool
}
