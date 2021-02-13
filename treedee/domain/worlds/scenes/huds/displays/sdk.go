package displays

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a display builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithIndex(index uint) Builder
	WithViewport(viewport viewports.Viewport) Builder
	WithCamera(cam cameras.Camera) Builder
	WithMaterial(mat materials.Material) Builder
	Now() (Display, error)
}

// Display represents a display
type Display interface {
	ID() *uuid.UUID
	Index() uint
	Viewport() viewports.Viewport
	Camera() cameras.Camera
	HasMaterial() bool
	Material() materials.Material
}
