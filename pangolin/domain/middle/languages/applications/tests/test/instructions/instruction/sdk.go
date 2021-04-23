package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithLanguage(lang standard_instruction.Instruction) Builder
	WithTest(test test_instruction.Instruction) Builder
	Now() (Instruction, error)
}

// Instruction represents a test instruction
type Instruction interface {
	IsLanguage() bool
	Language() standard_instruction.Instruction
	IsTest() bool
	Test() test_instruction.Instruction
}
