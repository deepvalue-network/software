package parsers

type variableSection struct {
	declarations []VariableDeclaration
}

func createVariableSection(declarations []VariableDeclaration) VariableSection {
	out := variableSection{
		declarations: declarations,
	}

	return &out
}

// Declarations returns the declarations
func (obj *variableSection) Declarations() []VariableDeclaration {
	return obj.declarations
}
