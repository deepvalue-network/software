package instruction

import (
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
)

type instruction struct {
	lang standard_instruction.CommonInstruction
	test test_instruction.Instruction
}

func createInstructionWithLanguage(
	lang standard_instruction.CommonInstruction,
) Instruction {
	return createInstructionInternally(lang, nil)
}

func createInstructionWithTest(
	test test_instruction.Instruction,
) Instruction {
	return createInstructionInternally(nil, test)
}

func createInstructionInternally(
	lang standard_instruction.CommonInstruction,
	test test_instruction.Instruction,
) Instruction {
	out := instruction{
		lang: lang,
		test: test,
	}

	return &out
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *instruction) IsLanguage() bool {
	return obj.lang != nil
}

// Language returns the language, if any
func (obj *instruction) Language() standard_instruction.CommonInstruction {
	return obj.lang
}

// IsTest returns true if there is a test, false otherwise
func (obj *instruction) IsTest() bool {
	return obj.test != nil
}

// Test returns the test, if any
func (obj *instruction) Test() test_instruction.Instruction {
	return obj.test
}
