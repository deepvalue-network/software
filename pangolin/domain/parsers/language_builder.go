package parsers

import "errors"

type languageBuilder struct {
	app LanguageApplication
	def LanguageDefinition
}

func createLanguageBuilder() LanguageBuilder {
	out := languageBuilder{
		app: nil,
		def: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder()
}

// WithApplication adds an application to the builder
func (app *languageBuilder) WithApplication(application LanguageApplication) LanguageBuilder {
	app.app = application
	return app
}

// WithDefinition adds a definition to the builder
func (app *languageBuilder) WithDefinition(definition LanguageDefinition) LanguageBuilder {
	app.def = definition
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
	if app.app != nil {
		return createLanguageWithApplication(app.app), nil
	}

	if app.def != nil {
		return createLanguageWithDefinition(app.def), nil
	}

	return nil, errors.New("the Language is invalid")
}
