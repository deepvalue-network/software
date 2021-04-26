package parsers

import "errors"

type testCommandInstructionBuilder struct {
	ins    TestInstruction
	scopes Scopes
}

func createTestCommandInstructionBuilder() TestCommandInstructionBuilder {
	out := testCommandInstructionBuilder{
		ins:    nil,
		scopes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testCommandInstructionBuilder) Create() TestCommandInstructionBuilder {
	return createTestCommandInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *testCommandInstructionBuilder) WithInstruction(ins TestInstruction) TestCommandInstructionBuilder {
	app.ins = ins
	return app
}

// WithScopes add scopes to the builder
func (app *testCommandInstructionBuilder) WithScopes(scopes Scopes) TestCommandInstructionBuilder {
	app.scopes = scopes
	return app
}

// Now builds a new TestCommandInstruction instance
func (app *testCommandInstructionBuilder) Now() (TestCommandInstruction, error) {
	if app.ins == nil {
		return nil, errors.New("the TestInstruction is mandatory in order to build a TestCommandInstruction instance")
	}

	if app.scopes != nil {
		return createTestCommandInstructionWithScopes(app.ins, app.scopes), nil
	}

	return createTestCommandInstruction(app.ins), nil
}
