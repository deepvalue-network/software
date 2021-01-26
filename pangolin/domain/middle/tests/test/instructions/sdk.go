package instructions

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/tests/test/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	instructionAdapter instruction.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(instructionAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the instructions adapter
type Adapter interface {
	ToInstructions(testInstructions []parsers.TestInstruction) (Instructions, error)
}

// Builder represents the instructions builder
type Builder interface {
	Create() Builder
	WithList(lst []instruction.Instruction) Builder
	WithMap(mp map[string]instruction.Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	All() []instruction.Instruction
}
