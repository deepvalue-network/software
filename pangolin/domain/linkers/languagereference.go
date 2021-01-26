package linkers

type languageReference struct {
	language Language
	input    string
	output   string
}

func createLanguageReference(
	language Language,
	input string,
	output string,
) LanguageReference {
	out := languageReference{
		language: language,
		input:    input,
		output:   output,
	}

	return &out
}

// Language returns the language
func (obj *languageReference) Language() Language {
	return obj.language
}

// Input returns the input
func (obj *languageReference) Input() string {
	return obj.input
}

// Output returns the output
func (obj *languageReference) Output() string {
	return obj.output
}
