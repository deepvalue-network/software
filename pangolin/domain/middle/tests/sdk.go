package tests

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/tests/test"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	testAdapter test.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(testAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents tests builder
type Adapter interface {
	ToTests(section parsers.TestSection) (Tests, error)
}

// Builder represents the tests builder
type Builder interface {
	Create() Builder
	WithList(lst []test.Test) Builder
	WithMap(mp map[string]test.Test) Builder
	Now() (Tests, error)
}

// Tests represents tests
type Tests interface {
	All() []test.Test
}
