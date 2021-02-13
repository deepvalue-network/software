package scenes

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the scene builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithIndex(index uint) Builder
	WithHud(hud huds.Hud) Builder
	WithNodes(nodes []nodes.Node) Builder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	ID() *uuid.UUID
	Index() uint
	Hud() huds.Hud
	Nodes() []nodes.Node
}
