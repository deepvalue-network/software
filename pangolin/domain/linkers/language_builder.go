package linkers

import "errors"

type languageBuilder struct {
	ref LanguageReference
	app LanguageApplication
}

func createLanguageBuilder() LanguageBuilder {
	out := languageBuilder{
		ref: nil,
		app: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder()
}

// WithReference adds a reference to the builder
func (app *languageBuilder) WithReference(ref LanguageReference) LanguageBuilder {
	app.ref = ref
	return app
}

// WithApplication adds an application to the builder
func (app *languageBuilder) WithApplication(appli LanguageApplication) LanguageBuilder {
	app.app = appli
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
	if app.ref != nil {
		return createLanguageWithReference(app.ref), nil
	}

	if app.app != nil {
		return createLanguageWithLanguageApplication(app.app), nil
	}

	return nil, errors.New("the Language is invalid")
}
