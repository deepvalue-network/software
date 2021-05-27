package labels

type label struct {
	name         string
	variable     string
	instructions []Instruction
}

func createLabel(
	name string,
	variable string,
	instructions []Instruction,
) Label {
	out := label{
		name:         name,
		variable:     variable,
		instructions: instructions,
	}

	return &out
}

// Name returns the name
func (obj *label) Name() string {
	return obj.name
}

// Variable returns the variable
func (obj *label) Variable() string {
	return obj.variable
}

// Instructions returns the instructions
func (obj *label) Instructions() []Instruction {
	return obj.instructions
}
