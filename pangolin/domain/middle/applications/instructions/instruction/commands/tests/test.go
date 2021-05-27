package tests

type test struct {
	name         string
	variable     string
	instructions []Instruction
}

func createTest(
	name string,
	variable string,
	instructions []Instruction,
) Test {
	out := test{
		name:         name,
		variable:     variable,
		instructions: instructions,
	}

	return &out
}

// Name returns the name
func (obj *test) Name() string {
	return obj.name
}

// Variable returns the variable
func (obj *test) Variable() string {
	return obj.variable
}

// Instructions returns the instructions
func (obj *test) Instructions() []Instruction {
	return obj.instructions
}
