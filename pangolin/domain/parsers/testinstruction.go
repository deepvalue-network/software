package parsers

type testInstruction struct {
	isStart  bool
	isStop   bool
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionWithStart() TestInstruction {
	return createTestInstructionInternally(true, false, nil, nil)
}

func createTestInstructionWithStop() TestInstruction {
	return createTestInstructionInternally(false, true, nil, nil)
}

func createTestInstructionWithInstruction(ins Instruction) TestInstruction {
	return createTestInstructionInternally(false, false, ins, nil)
}

func createTestInstructionWithReadFile(readFile ReadFile) TestInstruction {
	return createTestInstructionInternally(false, false, nil, readFile)
}

func createTestInstructionInternally(isStart bool, isStop bool, ins Instruction, readFile ReadFile) TestInstruction {
	out := testInstruction{
		isStart:  isStart,
		isStop:   isStop,
		readFile: readFile,
		ins:      ins,
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

// IsReadFile returns true if there is a readFile, false otherwise
func (obj *testInstruction) IsReadFile() bool {
	return obj.readFile != nil
}

// ReadFile returns the readFile, if any
func (obj *testInstruction) ReadFile() ReadFile {
	return obj.readFile
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *testInstruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *testInstruction) Instruction() Instruction {
	return obj.ins
}
