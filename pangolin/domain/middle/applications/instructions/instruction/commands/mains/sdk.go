package mains

import (
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionAdapter := application_instruction.NewAdapter()
	builder := NewBuilder()
	instructionBuilder := NewInstructionBuilder()
	return createAdapter(instructionAdapter, builder, instructionBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToMain(parsed parsers.MainCommand) (Main, error)
}

// Builder represents a main builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithInstructions(ins []Instruction) Builder
	Now() (Main, error)
}

// Main represents a main command
type Main interface {
	Variable() string
	Instructions() []Instruction
}

// InstructionBuilder represents a main instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithInstruction(ins application_instruction.Instruction) InstructionBuilder
	WithScopes(scopes []bool) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents a main instruction
type Instruction interface {
	Instruction() application_instruction.Instruction
	HasScopes() bool
	Scopes() []bool
}
