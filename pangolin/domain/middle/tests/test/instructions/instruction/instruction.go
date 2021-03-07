package instruction

import ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"

type instruction struct {
	assert   Assert
	readFile ReadFile
	ins      ins.Instruction
}

func createInstructionWithAssert(assert Assert) Instruction {
	return createInstructionInternally(assert, nil, nil)
}

func createInstructionWithReadFile(readFile ReadFile) Instruction {
	return createInstructionInternally(nil, readFile, nil)
}

func createInstructionWithInstruction(ins ins.Instruction) Instruction {
	return createInstructionInternally(nil, nil, ins)
}

func createInstructionInternally(
	assert Assert,
	readFile ReadFile,
	inst ins.Instruction,
) Instruction {
	out := instruction{
		assert:   assert,
		readFile: readFile,
		ins:      inst,
	}

	return &out
}

// IsAssert returns true if there is an assert, false otherwise
func (obj *instruction) IsAssert() bool {
	return obj.assert != nil
}

// Assert returns the assert, if any
func (obj *instruction) Assert() Assert {
	return obj.assert
}

// IsReadFile returns true if there is a readFile, false otherwise
func (obj *instruction) IsReadFile() bool {
	return obj.readFile != nil
}

// ReadFile returns the readFile, if any
func (obj *instruction) ReadFile() ReadFile {
	return obj.readFile
}

// IsInstruction returns true if there is an instruction
func (obj *instruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *instruction) Instruction() ins.Instruction {
	return obj.ins
}
