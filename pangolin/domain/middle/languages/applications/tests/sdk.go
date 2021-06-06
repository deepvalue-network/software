package tests

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	testAdapter := test.NewAdapter()
	builder := NewBuilder()
	return createAdapter(testAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a tests adapter
type Adapter interface {
	ToTests(parsed parsers.LanguageTestSection) (Tests, error)
}

// Builder represents the tests builder
type Builder interface {
	Create() Builder
	WithList(list []test.Test) Builder
	Now() (Tests, error)
}

// Tests represents tests
type Tests interface {
	All() []test.Test
}
