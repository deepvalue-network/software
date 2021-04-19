package parsers

type languageTestDeclaration struct {
	name string
	list []LanguageTestInstruction
}

func createLanguageTestDeclaration(
	name string,
	list []LanguageTestInstruction,
) LanguageTestDeclaration {
	out := languageTestDeclaration{
		name: name,
		list: list,
	}

	return &out
}

// Name returns the name
func (obj *languageTestDeclaration) Name() string {
	return obj.name
}

// Instructions returns the instructions
func (obj *languageTestDeclaration) Instructions() []LanguageTestInstruction {
	return obj.list
}
