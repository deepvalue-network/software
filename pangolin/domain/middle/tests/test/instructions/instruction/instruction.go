package instruction

import ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"

type instruction struct {
	isAssert bool
	readFile ReadFile
	ins      ins.Instruction
}

func createInstructionWithAssert() Instruction {
	return createInstructionInternally(true, nil, nil)
}

func createInstructionWithReadFile(readFile ReadFile) Instruction {
	return createInstructionInternally(false, readFile, nil)
}

func createInstructionWithInstruction(ins ins.Instruction) Instruction {
	return createInstructionInternally(false, nil, ins)
}

func createInstructionInternally(
	isAssert bool,
	readFile ReadFile,
	inst ins.Instruction,
) Instruction {
	out := instruction{
		isAssert: isAssert,
		readFile: readFile,
		ins:      inst,
	}

	return &out
}

// IsAssert returns true if there is an assert, false otherwise
func (obj *instruction) IsAssert() bool {
	return obj.isAssert
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
