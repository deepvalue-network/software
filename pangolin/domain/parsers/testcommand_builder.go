package parsers

import "errors"

type testCommandBuilder struct {
	variable string
	name     string
	ins      []TestCommandInstruction
}

func createTestCommandBuilder() TestCommandBuilder {
	out := testCommandBuilder{
		variable: "",
		name:     "",
		ins:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *testCommandBuilder) Create() TestCommandBuilder {
	return createTestCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *testCommandBuilder) WithVariable(variable string) TestCommandBuilder {
	app.variable = variable
	return app
}

// WithName adds a name to the builder
func (app *testCommandBuilder) WithName(name string) TestCommandBuilder {
	app.name = name
	return app
}

// WithInstructions adds instructions to the builder
func (app *testCommandBuilder) WithInstructions(ins []TestCommandInstruction) TestCommandBuilder {
	app.ins = ins
	return app
}

// Now builds a new TestCommand instance
func (app *testCommandBuilder) Now() (TestCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a TestCommand instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a TestCommand instance")
	}

	if app.ins != nil && len(app.ins) <= 0 {
		app.ins = nil
	}

	if app.ins == nil {
		return nil, errors.New("the instructions are mandatory in order to build a TestCommand instance")
	}

	return createTestCommand(app.variable, app.name, app.ins), nil
}
