package parsers

type testCommand struct {
	variable string
	name     string
	ins      []TestCommandInstruction
}

func createTestCommand(
	variable string,
	name string,
	ins []TestCommandInstruction,
) TestCommand {
	return createTestCommandInternally(variable, name, ins)
}

func createTestCommandInternally(
	variable string,
	name string,
	ins []TestCommandInstruction,
) TestCommand {
	out := testCommand{
		variable: variable,
		name:     name,
		ins:      ins,
	}

	return &out
}

// Variable returns the variable
func (obj *testCommand) Variable() string {
	return obj.variable
}

// Name returns the name
func (obj *testCommand) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *testCommand) Instructions() []TestCommandInstruction {
	return obj.ins
}
