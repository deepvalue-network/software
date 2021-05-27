package label

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionsAdapter := instructions.NewAdapter()
	builder := NewBuilder()
	return createAdapter(instructionsAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	instructionsBuilder := instructions.NewBuilder()
	return createBuilder(instructionsBuilder)
}

// Adapter represents the label adapter
type Adapter interface {
	ToLabel(declaration parsers.LabelDeclaration) (Label, error)
}

// Builder represents the label builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithInstructions(ins instructions.Instructions) Builder
	Now() (Label, error)
}

// Label represents a label
type Label interface {
	Name() string
	Instructions() instructions.Instructions
}
