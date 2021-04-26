package match

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a match adapter
type Adapter interface {
	ToMatch(parsed parsers.Match) (Match, error)
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
