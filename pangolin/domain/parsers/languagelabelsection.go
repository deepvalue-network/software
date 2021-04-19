package parsers

type languageLabelSection struct {
	declarations []LanguageLabelDeclaration
}

func createLanguageLabelSection(
	declarations []LanguageLabelDeclaration,
) LanguageLabelSection {
	out := languageLabelSection{
		declarations: declarations,
	}

	return &out
}

// Declarations returns the declarations
func (obj *languageLabelSection) Declarations() []LanguageLabelDeclaration {
	return obj.declarations
}
