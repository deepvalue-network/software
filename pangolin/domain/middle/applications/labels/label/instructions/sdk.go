package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionAdapter := instruction.NewAdapter()
	builder := NewBuilder()
	return createAdapter(instructionAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an instructions adapter
type Adapter interface {
	ToInstructions(parsed []parsers.LanguageLabelInstruction) (Instructions, error)
}

// Builder represents an instructions builder
type Builder interface {
	Create() Builder
	WithList(list []instruction.Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	All() []instruction.Instruction
}
