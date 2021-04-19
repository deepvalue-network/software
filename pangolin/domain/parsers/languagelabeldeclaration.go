package parsers

type languageLabelDeclaration struct {
	name string
	list []LanguageLabelInstruction
}

func createLanguageLabelDeclaration(
	name string,
	list []LanguageLabelInstruction,
) LanguageLabelDeclaration {
	out := languageLabelDeclaration{
		name: name,
		list: list,
	}

	return &out
}

// Name returns the name
func (obj *languageLabelDeclaration) Name() string {
	return obj.name
}

// Instructions returns the instructions list
func (obj *languageLabelDeclaration) Instructions() []LanguageLabelInstruction {
	return obj.list
}
