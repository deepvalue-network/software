package linkers

type languageReference struct {
	def   LanguageDefinition
	input string
}

func createLanguageReference(
	def LanguageDefinition,
	input string,
) LanguageReference {
	out := languageReference{
		def:   def,
		input: input,
	}

	return &out
}

// Definition returns the language definition
func (obj *languageReference) Definition() LanguageDefinition {
	return obj.def
}

// Input returns the input
func (obj *languageReference) Input() string {
	return obj.input
}
