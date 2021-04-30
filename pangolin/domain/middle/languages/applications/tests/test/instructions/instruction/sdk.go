package instruction

import (
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	languageAdapter := standard_instruction.NewAdapter()
	testAdapter := test_instruction.NewAdapter()
	builder := NewBuilder()
	return createAdapter(languageAdapter, testAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an instruction adapter
type Adapter interface {
	ToInstruction(parsed parsers.LanguageTestInstruction) (Instruction, error)
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
