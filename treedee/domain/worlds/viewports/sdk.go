package viewports

import "github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a viewport builder
type Builder interface {
	Create() Builder
	WithRectangle(rect ints.Rectangle) Builder
	WithVariable(variable string) Builder
	Now() (Viewport, error)
}

// Viewport represents a viewport
type Viewport interface {
	Rectangle() ints.Rectangle
	Variable() string
	IsContained(dim ints.Vec2) bool
}
