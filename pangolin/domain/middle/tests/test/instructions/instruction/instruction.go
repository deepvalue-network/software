package instruction

import ins "github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"

type instruction struct {
	isStart bool
	isStop  bool
	ins     ins.Instruction
}

func createInstructionWithStart() Instruction {
	return createInstructionInternally(true, false, nil)
}

func createInstructionWithStop() Instruction {
	return createInstructionInternally(false, true, nil)
}

func createInstructionWithInstruction(ins ins.Instruction) Instruction {
	return createInstructionInternally(false, false, ins)
}

func createInstructionInternally(
	isStart bool,
	isStop bool,
	inst ins.Instruction,
) Instruction {
	out := instruction{
		isStart: isStart,
		isStop:  isStop,
		ins:     inst,
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

// IsInstruction returns true if there is an instruction
func (obj *instruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *instruction) Instruction() ins.Instruction {
	return obj.ins
}
