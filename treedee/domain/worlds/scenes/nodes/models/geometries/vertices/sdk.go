package vertices

import "github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a vertex builder
type Builder interface {
	Create() Builder
	WithSpace(pos fl32.Vec3) Builder
	WithTexture(tex fl32.Vec2) Builder
	Now() (Vertex, error)
}

// Vertex represents a vertex
type Vertex interface {
	Space() fl32.Vec3
	Texture() fl32.Vec2
}
