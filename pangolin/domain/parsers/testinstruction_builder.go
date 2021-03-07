package parsers

import "errors"

type testInstructionBuilder struct {
	assert   Assert
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionBuilder() TestInstructionBuilder {
	out := testInstructionBuilder{
		assert:   nil,
		readFile: nil,
		ins:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *testInstructionBuilder) Create() TestInstructionBuilder {
	return createTestInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *testInstructionBuilder) WithInstruction(ins Instruction) TestInstructionBuilder {
	app.ins = ins
	return app
}

// WithReadFile adds a readFile to the builder
func (app *testInstructionBuilder) WithReadFile(readFile ReadFile) TestInstructionBuilder {
	app.readFile = readFile
	return app
}

// WithAssert adds an assert to the builder
func (app *testInstructionBuilder) WithAssert(assert Assert) TestInstructionBuilder {
	app.assert = assert
	return app
}

// Now builds a new TestInstruction instance
func (app *testInstructionBuilder) Now() (TestInstruction, error) {
	if app.ins != nil {
		return createTestInstructionWithInstruction(app.ins), nil
	}

	if app.readFile != nil {
		return createTestInstructionWithReadFile(app.readFile), nil
	}

	if app.assert != nil {
		return createTestInstructionWithAssert(app.assert), nil
	}

	return nil, errors.New("the TestInstruction is invalid")
}
