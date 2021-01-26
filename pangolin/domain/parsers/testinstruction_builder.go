package parsers

import "errors"

type testInstructionBuilder struct {
	isStart  bool
	isStop   bool
	readFile ReadFile
	ins      Instruction
}

func createTestInstructionBuilder() TestInstructionBuilder {
	out := testInstructionBuilder{
		isStart:  false,
		isStop:   false,
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

// IsStart flags the instruction as a start
func (app *testInstructionBuilder) IsStart() TestInstructionBuilder {
	app.isStart = true
	return app
}

// IsStop flags the instruction as a stop
func (app *testInstructionBuilder) IsStop() TestInstructionBuilder {
	app.isStop = true
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

	if app.isStop {
		return createTestInstructionWithStop(), nil
	}

	if app.isStart {
		return createTestInstructionWithStart(), nil
	}

	return nil, errors.New("the TestInstruction is invalid")
}
