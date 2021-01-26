package parsers

type labelInstruction struct {
	isRet       bool
	instruction Instruction
}

func createLabelInstructionWithRet() LabelInstruction {
	return createLabelInstructionInternally(true, nil)
}

func createLabelInstructionWithInstruction(instruction Instruction) LabelInstruction {
	return createLabelInstructionInternally(false, instruction)
}

func createLabelInstructionInternally(isRet bool, instruction Instruction) LabelInstruction {
	out := labelInstruction{
		isRet:       isRet,
		instruction: instruction,
	}

	return &out
}

// IsRet returns true if the label instruction is return, false otherwise
func (obj *labelInstruction) IsRet() bool {
	return obj.isRet
}

// IsInstruction returns true if the label instruction is an instruction, false otherwise
func (obj *labelInstruction) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *labelInstruction) Instruction() Instruction {
	return obj.instruction
}
