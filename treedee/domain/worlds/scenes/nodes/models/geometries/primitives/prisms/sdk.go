package prisms

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

// Adapter represents a prism adapter
type Adapter interface {
	ToVertices(prism Prism) ([]vertices.Vertex, error)
}

// Builder represents a prism builder
type Builder interface {
	Create() Builder
	WithHeight(height float32) Builder
	WithBase(base []fl32.Vec2) Builder
	IsTopCorner() Builder
	Now() (Prism, error)
}

// Prism represents a prism
type Prism interface {
	Height() float32
	Base() []fl32.Vec2
	IsTopCorner() bool
}
