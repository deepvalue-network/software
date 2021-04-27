package label

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	instructionsAdapter instructions.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(instructionsAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a label adapter
type Adapter interface {
	ToLabel(parsed parsers.LanguageLabelDeclaration) (Label, error)
}

// Builder represents a label builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	Now() (Label, error)
}

// Label represents a label
type Label interface {
	Name() string
	Instructions() instructions.Instructions
}
