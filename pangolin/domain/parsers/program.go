package parsers

type program struct {
	testable Testable
	language LanguageApplication
}

func createProgramWithTestable(
	testable Testable,
) Program {
	return createProgramInternally(testable, nil)
}

func createProgramWithLanguage(
	language LanguageApplication,
) Program {
	return createProgramInternally(nil, language)
}

func createProgramInternally(
	testable Testable,
	language LanguageApplication,
) Program {
	out := program{
		testable: testable,
		language: language,
	}

	return &out
}

// IsTestable returns true if the program is testable, false otherwise
func (obj *program) IsTestable() bool {
	return obj.testable != nil
}

// Testable returns the testable, if any
func (obj *program) Testable() Testable {
	return obj.testable
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *program) IsLanguage() bool {
	return obj.language != nil
}

// Language returns the language, if any
func (obj *program) Language() LanguageApplication {
	return obj.language
}
