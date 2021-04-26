package tests

import (
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	testInstructionAdapter test_instruction.Adapter,
) Adapter {
	builder := NewBuilder()
	instructionBuilder := NewInstructionBuilder()
	return createAdapter(testInstructionAdapter, builder, instructionBuilder)
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
	ToTest(parsed parsers.TestCommand) (Test, error)
}

// Builder represents a test builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVariable(variable string) Builder
	WithInstructions(ins []Instruction) Builder
	Now() (Test, error)
}

// Test represents a test command
type Test interface {
	Name() string
	Variable() string
	Instructions() []Instruction
}

// InstructionBuilder represents a test instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithInstruction(ins test_instruction.Instruction) InstructionBuilder
	WithScopes(scopes []bool) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents a test instruction
type Instruction interface {
	Instruction() test_instruction.Instruction
	HasScopes() bool
	Scopes() []bool
}
