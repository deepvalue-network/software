package layers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/alphas"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithIndex(index uint) Builder
	WithAlpha(alpha alphas.Alpha) Builder
	WithTexture(tex textures.Texture) Builder
	Now() (Layer, error)
}

// Layer represents layer of textures
type Layer interface {
	ID() *uuid.UUID
	Index() uint
	Alpha() alphas.Alpha
	Texture() textures.Texture
}
