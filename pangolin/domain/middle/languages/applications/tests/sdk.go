package tests

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
