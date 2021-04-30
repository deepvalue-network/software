package linkers

import "errors"

type languageReferenceBuilder struct {
	def   LanguageDefinition
	input string
}

func createLanguageReferenceBuilder() LanguageReferenceBuilder {
	out := languageReferenceBuilder{
		def:   nil,
		input: "",
	}

	return &out
}

// Create initializes the builder
func (app *languageReferenceBuilder) Create() LanguageReferenceBuilder {
	return createLanguageReferenceBuilder()
}

// WithDefinition adds a language definition to the builder
func (app *languageReferenceBuilder) WithDefinition(def LanguageDefinition) LanguageReferenceBuilder {
	app.def = def
	return app
}

// WithInputVariable adds an input variable to the builder
func (app *languageReferenceBuilder) WithInputVariable(input string) LanguageReferenceBuilder {
	app.input = input
	return app
}

// Now builds a new LanguageReference instance
func (app *languageReferenceBuilder) Now() (LanguageReference, error) {
	if app.def == nil {
		return nil, errors.New("the languageDefinition is mandatory in order to build a LanguageReference instance")
	}

	if app.input == "" {
		return nil, errors.New("the input variable is mandatory in order to build a LanguageReference instance")
	}

	return createLanguageReference(app.def, app.input), nil
}
