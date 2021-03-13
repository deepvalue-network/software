package format

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a format builder
type Builder interface {
	Create() Builder
	WithResults(results string) Builder
	WithPattern(pattern string) Builder
	WithFirst(first string) Builder
	WithSecond(second string) Builder
	Now() (Format, error)
}

// Format represents a format
type Format interface {
	Results() string
	Pattern() string
	First() string
	Second() string
}
