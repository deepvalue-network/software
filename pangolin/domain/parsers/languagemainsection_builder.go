package parsers

import "errors"

type languageMainSectionBuilder struct {
	instructions []LanguageInstruction
}

func createLanguageMainSectionBuilder() LanguageMainSectionBuilder {
	out := languageMainSectionBuilder{
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageMainSectionBuilder) Create() LanguageMainSectionBuilder {
	return createLanguageMainSectionBuilder()
}

// WithInstructions add instructions to the builder
func (app *languageMainSectionBuilder) WithInstructions(instructions []LanguageInstruction) LanguageMainSectionBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new LanguageMainSection instance
func (app *languageMainSectionBuilder) Now() (LanguageMainSection, error) {
	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("there must be at least 1 LanguageInstruction instance in order to build a LanguageMainSection instance")
	}

	return createLanguageMainSection(app.instructions), nil
}
