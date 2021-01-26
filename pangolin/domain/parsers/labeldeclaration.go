package parsers

type labelDeclaration struct {
	name         string
	instructions []LabelInstruction
}

func createLabelDeclaration(name string, instructions []LabelInstruction) LabelDeclaration {
	out := labelDeclaration{
		name:         name,
		instructions: instructions,
	}

	return &out
}

// Name returns the name
func (obj *labelDeclaration) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *labelDeclaration) Instructions() []LabelInstruction {
	return obj.instructions
}
