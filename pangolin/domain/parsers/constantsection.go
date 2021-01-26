package parsers

type constantSection struct {
	declarations []ConstantDeclaration
}

func createConstantSection(declarations []ConstantDeclaration) ConstantSection {
	out := constantSection{
		declarations: declarations,
	}

	return &out
}

// Declarations returns the declarations
func (obj *constantSection) Declarations() []ConstantDeclaration {
	return obj.declarations
}
