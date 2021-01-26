package parsers

type mainSection struct {
	instructions []Instruction
}

func createMainSection(instructions []Instruction) MainSection {
	out := mainSection{
		instructions: instructions,
	}

	return &out
}

// Instructions return the instructions
func (obj *mainSection) Instructions() []Instruction {
	return obj.instructions
}
