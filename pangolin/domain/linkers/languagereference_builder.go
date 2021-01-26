package linkers

import "errors"

type languageReferenceBuilder struct {
	language Language
	input    string
	output   string
}

func createLanguageReferenceBuilder() LanguageReferenceBuilder {
	out := languageReferenceBuilder{
		language: nil,
		input:    "",
		output:   "",
	}

	return &out
}

// Create initializes the builder
func (app *languageReferenceBuilder) Create() LanguageReferenceBuilder {
	return createLanguageReferenceBuilder()
}

// WithLanguage adds a language to the builder
func (app *languageReferenceBuilder) WithLanguage(language Language) LanguageReferenceBuilder {
	app.language = language
	return app
}

// WithInputVariable adds an input variable to the builder
func (app *languageReferenceBuilder) WithInputVariable(input string) LanguageReferenceBuilder {
	app.input = input
	return app
}

// WithOutputVariable adds an output variable to the builder
func (app *languageReferenceBuilder) WithOutputVariable(output string) LanguageReferenceBuilder {
	app.output = output
	return app
}

// Now builds a new LanguageReference instance
func (app *languageReferenceBuilder) Now() (LanguageReference, error) {
	if app.language == nil {
		return nil, errors.New("the language is mandatory in order to build a LanguageReference instance")
	}

	if app.input == "" {
		return nil, errors.New("the input variable is mandatory in order to build a LanguageReference instance")
	}

	if app.output == "" {
		return nil, errors.New("the output variable is mandatory in order to build a LanguageReference instance")
	}

	return createLanguageReference(app.language, app.input, app.output), nil
}
