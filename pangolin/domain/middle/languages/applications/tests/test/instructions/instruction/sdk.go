package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/tests/test/instructions/instruction"
)

// Instruction represents a test instruction
type Instruction interface {
	IsLanguage() bool
	Language() instruction.Instruction
	IsTest() bool
	Test() test_instruction.Instruction
}
