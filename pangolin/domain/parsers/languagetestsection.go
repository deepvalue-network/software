package parsers

type languageTestSection struct {
	declarations []LanguageTestDeclaration
}

func createLanguageTestSection(
	declarations []LanguageTestDeclaration,
) LanguageTestSection {
	out := languageTestSection{
		declarations: declarations,
	}

	return &out
}

// Declarations returns the declarations
func (obj *languageTestSection) Declarations() []LanguageTestDeclaration {
	return obj.declarations
}
