package labels

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToLabel(parsed parsers.LabelCommand) (Label, error)
}

// Builder represents a label builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVariable(variable string) Builder
	WithInstructions(ins []Instruction) Builder
	Now() (Label, error)
}

// Label represents a label command
type Label interface {
	Name() string
	Variable() string
	Instructions() []Instruction
}

// InstructionBuilder represents a label instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithInstruction(ins label_instruction.Instruction) InstructionBuilder
	WithScopes(scopes []bool) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents a label instruction
type Instruction interface {
	Instruction() label_instruction.Instruction
	HasScopes() bool
	Scopes() []bool
}
