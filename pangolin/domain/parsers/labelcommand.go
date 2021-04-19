package parsers

type labelCommand struct {
	variable string
	name     string
	ins      []LabelCommandInstruction
}

func createLabelCommand(
	variable string,
	name string,
	ins []LabelCommandInstruction,
) LabelCommand {
	return createLabelCommandInternally(variable, name, ins)
}

func createLabelCommandInternally(
	variable string,
	name string,
	ins []LabelCommandInstruction,
) LabelCommand {
	out := labelCommand{
		variable: variable,
		name:     name,
		ins:      ins,
	}

	return &out
}

// Variable returns the variable
func (obj *labelCommand) Variable() string {
	return obj.variable
}

// Name returns the name
func (obj *labelCommand) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *labelCommand) Instructions() []LabelCommandInstruction {
	return obj.ins
}
