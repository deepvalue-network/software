package colors

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a color builder
type Builder interface {
	Create() Builder
	WithRed(red uint8) Builder
	WithGreen(green uint8) Builder
	WithBlue(blue uint8) Builder
	Now() Color
}

// Color represents a color
type Color interface {
	Red() uint8
	Green() uint8
	Blue() uint8
}
