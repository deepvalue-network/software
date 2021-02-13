package ints

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Vec2 represents a vector with 2 elements
type Vec2 [2]int

// Builder represents a rectangle builder
type Builder interface {
	Create() Builder
	WithPosition(pos Vec2) Builder
	WithDimension(dim Vec2) Builder
	Now() (Rectangle, error)
}

// Rectangle represents a rectangle of integers
type Rectangle interface {
	Position() Vec2
	Dimension() Vec2
	String() string
}
