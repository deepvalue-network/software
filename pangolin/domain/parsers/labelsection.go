package parsers

type labelSection struct {
	declarations []LabelDeclaration
}

func createLabelSection(declarations []LabelDeclaration) LabelSection {
	out := labelSection{
		declarations: declarations,
	}

	return &out
}

// Declarations return labelDeclarations
func (obj *labelSection) Declarations() []LabelDeclaration {
	return obj.declarations
}
