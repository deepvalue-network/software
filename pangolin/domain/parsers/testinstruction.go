package parsers

type testInstruction struct {
	isStart bool
	isStop  bool
	ins     Instruction
}

func createTestInstructionWithStart() TestInstruction {
	return createTestInstructionInternally(true, false, nil)
}

func createTestInstructionWithStop() TestInstruction {
	return createTestInstructionInternally(false, true, nil)
}

func createTestInstructionWithInstruction(ins Instruction) TestInstruction {
	return createTestInstructionInternally(false, false, ins)
}

func createTestInstructionInternally(isStart bool, isStop bool, ins Instruction) TestInstruction {
	out := testInstruction{
		isStart: isStart,
		isStop:  isStop,
		ins:     ins,
	}

	return &out
}

// IsStart returns true if the instruction is start, false otherwise
func (obj *testInstruction) IsStart() bool {
	return obj.isStart
}

// IsStop returns true if the instruction is stop, false otherwise
func (obj *testInstruction) IsStop() bool {
	return obj.isStop
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *testInstruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *testInstruction) Instruction() Instruction {
	return obj.ins
}
