package match

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a matcher builder
type Builder interface {
	Create() Builder
	WithInput(input string) Builder
	WithPattern(pattern string) Builder
	Now() (Match, error)
}

// Match represents a match
type Match interface {
	Input() string
	HasPattern() bool
	Pattern() string
}
