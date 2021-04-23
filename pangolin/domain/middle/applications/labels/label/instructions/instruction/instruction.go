package instruction

import (
	ins "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
)

type instruction struct {
	isRet bool
	ins   ins.Instruction
}

func createInstructionWithReturn() Instruction {
	return createInstructionInternally(true, nil)
}

func createInstructionWithInstruction(ins ins.Instruction) Instruction {
	return createInstructionInternally(false, ins)
}

func createInstructionInternally(isRet bool, ins ins.Instruction) Instruction {
	out := instruction{
		isRet: isRet,
		ins:   ins,
	}

	return &out
}

// IsRet returns true if the instruction is return, false otherwise
func (obj *instruction) IsRet() bool {
	return obj.isRet
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *instruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *instruction) Instruction() ins.Instruction {
	return obj.ins
}
