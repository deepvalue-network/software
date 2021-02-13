package worlds

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes"
	uuid "github.com/satori/go.uuid"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a world builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithCurrentSceneIndex(currentSceneIndex uint) Builder
	WithScenes(scenes []scenes.Scene) Builder
	Now() (World, error)
}

// World represents a world
type World interface {
	ID() *uuid.UUID
	CurrentSceneIndex() uint
	Scenes() []scenes.Scene
}
