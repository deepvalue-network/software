package parsers

import "errors"

type testInstructionBuilder struct {
	isAssert bool
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionBuilder() TestInstructionBuilder {
	out := testInstructionBuilder{
		isAssert: false,
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

// IsAssert adds an assert to the builder
func (app *testInstructionBuilder) IsAssert() TestInstructionBuilder {
	app.isAssert = true
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

	if app.isAssert {
		return createTestInstructionWithAssert(), nil
	}

	return nil, errors.New("the TestInstruction is invalid")
}
