package parsers

type languageMainSection struct {
	list []LanguageInstruction
}

func createLanguageMainSection(
	list []LanguageInstruction,
) LanguageMainSection {
	out := languageMainSection{
		list: list,
	}

	return &out
}

// Declarations returns the instructions
func (obj *languageMainSection) Instructions() []LanguageInstruction {
	return obj.list
}
