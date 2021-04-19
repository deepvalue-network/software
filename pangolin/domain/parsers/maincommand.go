package parsers

type mainCommand struct {
	variable string
	ins      []MainCommandInstruction
}

func createMainCommand(
	variable string,
	ins []MainCommandInstruction,
) MainCommand {
	return createMainCommandInternally(variable, ins)
}

func createMainCommandInternally(
	variable string,
	ins []MainCommandInstruction,
) MainCommand {
	out := mainCommand{
		variable: variable,
		ins:      ins,
	}

	return &out
}

// Variable returns the variable
func (obj *mainCommand) Variable() string {
	return obj.variable
}

// Instructions returns the instructions
func (obj *mainCommand) Instructions() []MainCommandInstruction {
	return obj.ins
}
