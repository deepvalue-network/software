package test

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionsAdapter := instructions.NewAdapter()
	builder := NewBuilder()
	return createAdapter(instructionsAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a test adapter
type Adapter interface {
	ToTest(declaration parsers.TestDeclaration) (Test, error)
}

// Builder represents a test builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithInstructions(ins instructions.Instructions) Builder
	Now() (Test, error)
}

// Test represents a test
type Test interface {
	Name() string
	Instructions() instructions.Instructions
}
