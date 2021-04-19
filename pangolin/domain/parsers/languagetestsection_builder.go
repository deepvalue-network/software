package parsers

import "errors"

type languageTestSectionBuilder struct {
	declarations []LanguageTestDeclaration
}

func createLanguageTestSectionBuilder() LanguageTestSectionBuilder {
	out := languageTestSectionBuilder{
		declarations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageTestSectionBuilder) Create() LanguageTestSectionBuilder {
	return createLanguageTestSectionBuilder()
}

// WithDeclarations add declarations to the builder
func (app *languageTestSectionBuilder) WithDeclarations(declarations []LanguageTestDeclaration) LanguageTestSectionBuilder {
	app.declarations = declarations
	return app
}

// Now builds a new LanguageTestSection instance
func (app *languageTestSectionBuilder) Now() (LanguageTestSection, error) {
	if app.declarations != nil && len(app.declarations) <= 0 {
		app.declarations = nil
	}

	if app.declarations == nil {
		return nil, errors.New("there must be at least 1 LanguageTestDeclaration instance in order to build a LanguageTestSection instance")
	}

	return createLanguageTestSection(app.declarations), nil
}
