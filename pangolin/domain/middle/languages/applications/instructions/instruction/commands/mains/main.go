package mains

type main struct {
	variable     string
	instructions []Instruction
}

func createMain(
	variable string,
	instructions []Instruction,
) Main {
	out := main{
		variable:     variable,
		instructions: instructions,
	}

	return &out
}

// Variable returns the variable
func (obj *main) Variable() string {
	return obj.variable
}

// Instructions returns the instructions
func (obj *main) Instructions() []Instruction {
	return obj.instructions
}
