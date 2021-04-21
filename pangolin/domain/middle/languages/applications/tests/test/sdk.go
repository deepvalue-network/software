package test

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
