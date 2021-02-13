package spheres

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

// Adapter represents a sphere adapter
type Adapter interface {
	ToVertices(sphere Sphere) ([]vertices.Vertex, error)
}

// Builder represents a sphere builder
type Builder interface {
	Create() Builder
	WithPoints(points uint) Builder
	WithRadius(radius float32) Builder
	WithCenter(center fl32.Vec3) Builder
	Now() (Sphere, error)
}

// Sphere represents a sphere
type Sphere interface {
	Points() uint
	Radius() float32
	Center() fl32.Vec3
}
