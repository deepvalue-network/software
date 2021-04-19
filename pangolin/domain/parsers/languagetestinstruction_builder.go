package parsers

import "errors"

type languageTestInstructionBuilder struct {
	lang LanguageInstruction
	test TestInstruction
}

func createLanguageTestInstructionBuilder() LanguageTestInstructionBuilder {
	out := languageTestInstructionBuilder{
		lang: nil,
		test: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageTestInstructionBuilder) Create() LanguageTestInstructionBuilder {
	return createLanguageTestInstructionBuilder()
}

// WithLanguageInstruction adds a language instruction to the builder
func (app *languageTestInstructionBuilder) WithLanguageInstruction(languageIns LanguageInstruction) LanguageTestInstructionBuilder {
	app.lang = languageIns
	return app
}

// WithTestInstruction adds a test instruction to the builder
func (app *languageTestInstructionBuilder) WithTestInstruction(testIns TestInstruction) LanguageTestInstructionBuilder {
	app.test = testIns
	return app
}

// Now builds a new LanguageTestInstruction instance
func (app *languageTestInstructionBuilder) Now() (LanguageTestInstruction, error) {
	if app.lang != nil {
		return createLanguageTestInstructionWithLanguage(app.lang), nil
	}

	if app.test != nil {
		return createLanguageTestInstructionWithTest(app.test), nil
	}

	return nil, errors.New("the LanguageTestInstruction instance is invalid")
}
