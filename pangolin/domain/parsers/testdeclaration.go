package parsers

type testDeclaration struct {
	name         string
	instructions []TestInstruction
}

func createTestDeclaration(name string, instructions []TestInstruction) TestDeclaration {
	out := testDeclaration{
		name:         name,
		instructions: instructions,
	}

	return &out
}

// Name returns the name
func (obj *testDeclaration) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *testDeclaration) Instructions() []TestInstruction {
	return obj.instructions
}
