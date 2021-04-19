package parsers

import "errors"

type languageLabelSectionBuilder struct {
	declarations []LanguageLabelDeclaration
}

func createLanguageLabelSectionBuilder() LanguageLabelSectionBuilder {
	out := languageLabelSectionBuilder{
		declarations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageLabelSectionBuilder) Create() LanguageLabelSectionBuilder {
	return createLanguageLabelSectionBuilder()
}

// WithDeclarations add declarations to the builder
func (app *languageLabelSectionBuilder) WithDeclarations(declarations []LanguageLabelDeclaration) LanguageLabelSectionBuilder {
	app.declarations = declarations
	return app
}

// Now builds a new LanguageLabelSection instance
func (app *languageLabelSectionBuilder) Now() (LanguageLabelSection, error) {
	if app.declarations != nil && len(app.declarations) <= 0 {
		app.declarations = nil
	}

	if app.declarations == nil {
		return nil, errors.New("there must be at least 1 LanguageLabelDeclaration instance in order to build a LanguageLabelSection instance")
	}

	return createLanguageLabelSection(app.declarations), nil
}
