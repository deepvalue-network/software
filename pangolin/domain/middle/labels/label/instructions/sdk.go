package instructions

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/labels/label/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
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

// Adapter represents the label instructions adapter
type Adapter interface {
	ToInstructions(instructions []parsers.LabelInstruction) (Instructions, error)
}

// Builder represents the label instructions builder
type Builder interface {
	Create() Builder
	WithList(lst []instruction.Instruction) Builder
	WithMap(mp map[string]instruction.Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents label instructions
type Instructions interface {
	All() []instruction.Instruction
}
