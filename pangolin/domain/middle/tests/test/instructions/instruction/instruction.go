package instruction

import ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"

type instruction struct {
	isStart  bool
	isStop   bool
	readFile ReadFile
	ins      ins.Instruction
}

func createInstructionWithStart() Instruction {
	return createInstructionInternally(true, false, nil, nil)
}

func createInstructionWithStop() Instruction {
	return createInstructionInternally(false, true, nil, nil)
}

func createInstructionWithReadFile(readFile ReadFile) Instruction {
	return createInstructionInternally(false, false, readFile, nil)
}

func createInstructionWithInstruction(ins ins.Instruction) Instruction {
	return createInstructionInternally(false, false, nil, ins)
}

func createInstructionInternally(
	isStart bool,
	isStop bool,
	readFile ReadFile,
	inst ins.Instruction,
) Instruction {
	out := instruction{
		isStart:  isStart,
		isStop:   isStop,
		readFile: readFile,
		ins:      inst,
	}

	return &out
}

// IsStart returns true if the instruction is start
func (obj *instruction) IsStart() bool {
	return obj.isStart
}

// IsStop returns true if the instruction is stop
func (obj *instruction) IsStop() bool {
	return obj.isStop
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
