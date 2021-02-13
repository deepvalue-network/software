package cubes

import "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"

// Adapter represents a cube adapter
type Adapter interface {
	ToVertices(cube Cube) ([]vertices.Vertex, error)
}

// Builder represents a cube builder
type Builder interface {
	Create() Builder
	WithWidth(width float32) Builder
	WithHeight(height float32) Builder
	WithDepth(depth float32) Builder
	Now() (Cube, error)
}

// Cube represents a cube
type Cube interface {
	Width() float32
	Height() float32
	Depth() float32
}
