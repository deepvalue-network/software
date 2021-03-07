package parsers

type testInstruction struct {
	assert   Assert
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionWithInstruction(ins Instruction) TestInstruction {
	return createTestInstructionInternally(ins, nil, nil)
}

func createTestInstructionWithAssert(assert Assert) TestInstruction {
	return createTestInstructionInternally(nil, assert, nil)
}

func createTestInstructionWithReadFile(readFile ReadFile) TestInstruction {
	return createTestInstructionInternally(nil, nil, readFile)
}

func createTestInstructionInternally(ins Instruction, assert Assert, readFile ReadFile) TestInstruction {
	out := testInstruction{
		assert:   assert,
		readFile: readFile,
		ins:      ins,
	}

	return &out
}

// IsAssert returns true if there is an assert, false otherwise
func (obj *testInstruction) IsAssert() bool {
	return obj.assert != nil
}

// Assert returns the assert, if any
func (obj *testInstruction) Assert() Assert {
	return obj.assert
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
