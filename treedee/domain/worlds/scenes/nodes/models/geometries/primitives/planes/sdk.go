package planes

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

// Adapter represents a plane adapter
type Adapter interface {
	ToVertices(plane Plane) ([]vertices.Vertex, error)
}

// Builder represents a plane builder
type Builder interface {
	Create() Builder
	WithPosition(pos fl32.Vec2) Builder
	WithWidth(width float32) Builder
	WithHeight(height float32) Builder
	Now() (Plane, error)
}

// Plane represents a plane
type Plane interface {
	Position() fl32.Vec2
	Width() float32
	Height() float32
}
