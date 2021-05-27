package parsers

type testable struct {
	executable Executable
	language   LanguageDefinition
}

func createTestableWithExecutable(
	executable Executable,
) Testable {
	return createTestableInternally(executable, nil)
}

func createTestableWithLanguage(
	language LanguageDefinition,
) Testable {
	return createTestableInternally(nil, language)
}

func createTestableInternally(
	executable Executable,
	language LanguageDefinition,
) Testable {
	out := testable{
		executable: executable,
		language:   language,
	}

	return &out
}

// IsExecutable returns true if there is an executable, false otherwise
func (obj *testable) IsExecutable() bool {
	return obj.executable != nil
}

// Executable returns the executable, if any
func (obj *testable) Executable() Executable {
	return obj.executable
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *testable) IsLanguage() bool {
	return obj.language != nil
}

// Language returns the language, if any
func (obj *testable) Language() LanguageDefinition {
	return obj.language
}
