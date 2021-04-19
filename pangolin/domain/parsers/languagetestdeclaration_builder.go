package parsers

import "errors"

type languageTestDeclarationBuilder struct {
	name string
	list []LanguageTestInstruction
}

func createLanguageTestDeclarationBuilder() LanguageTestDeclarationBuilder {
	out := languageTestDeclarationBuilder{
		name: "",
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageTestDeclarationBuilder) Create() LanguageTestDeclarationBuilder {
	return createLanguageTestDeclarationBuilder()
}

// WithName adds a name to the builder
func (app *languageTestDeclarationBuilder) WithName(name string) LanguageTestDeclarationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *languageTestDeclarationBuilder) WithInstructions(instructions []LanguageTestInstruction) LanguageTestDeclarationBuilder {
	app.list = instructions
	return app
}

// Now builds a new LanguageTestDeclaration instance
func (app *languageTestDeclarationBuilder) Now() (LanguageTestDeclaration, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a LanguageTestDeclaration instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 LanguageTestInstruction in order to build a LanguageTestDeclaration instance")
	}

	return createLanguageTestDeclaration(app.name, app.list), nil
}
