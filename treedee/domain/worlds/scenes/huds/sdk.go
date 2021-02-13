package huds

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/nodes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an HUD builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithNodes(nodes []nodes.Node) Builder
	WithMaterial(mat materials.Material) Builder
	Now() (Hud, error)
}

// Hud represents the head-up display
type Hud interface {
	ID() *uuid.UUID
	HasNodes() bool
	Nodes() []nodes.Node
	HasMaterial() bool
	Material() materials.Material
}
