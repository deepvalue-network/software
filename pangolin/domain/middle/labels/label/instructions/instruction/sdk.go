package instruction

import (
	ins "github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionAdapter := ins.NewAdapter()
	builder := NewBuilder()
	return createAdapter(instructionAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the label instruction adapter
type Adapter interface {
	ToInstruction(instruction parsers.LabelInstruction) (Instruction, error)
}

// Builder represents a label instruction builder
type Builder interface {
	Create() Builder
	IsRet() Builder
	WithInstruction(ins ins.Instruction) Builder
	Now() (Instruction, error)
}

// Instruction represents a label instruction
type Instruction interface {
	IsRet() bool
	IsInstruction() bool
	Instruction() ins.Instruction
}
