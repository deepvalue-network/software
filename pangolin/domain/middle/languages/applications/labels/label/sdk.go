package label

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
