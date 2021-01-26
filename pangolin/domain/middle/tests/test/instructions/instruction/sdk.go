package instruction

import (
	ins "github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	instructionAdapter ins.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(builder, instructionAdapter)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an instruction adapter
type Adapter interface {
	ToInstruction(testInstruction parsers.TestInstruction) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	IsStart() Builder
	IsStop() Builder
	WithInstruction(ins ins.Instruction) Builder
	Now() (Instruction, error)
}

// Instruction represents a test instruction
type Instruction interface {
	IsStart() bool
	IsStop() bool
	IsInstruction() bool
	Instruction() ins.Instruction
}
