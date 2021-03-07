package parsers

type testInstruction struct {
	isAssert bool
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionWithInstruction(ins Instruction) TestInstruction {
	return createTestInstructionInternally(ins, false, nil)
}

func createTestInstructionWithAssert() TestInstruction {
	return createTestInstructionInternally(nil, true, nil)
}

func createTestInstructionWithReadFile(readFile ReadFile) TestInstruction {
	return createTestInstructionInternally(nil, false, readFile)
}

func createTestInstructionInternally(ins Instruction, isAssert bool, readFile ReadFile) TestInstruction {
	out := testInstruction{
		isAssert: isAssert,
		readFile: readFile,
		ins:      ins,
	}

	return &out
}

// IsAssert returns true if there is an assert, false otherwise
func (obj *testInstruction) IsAssert() bool {
	return obj.isAssert
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
