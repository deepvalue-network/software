package parsers

import "errors"

type languageLabelDeclarationBuilder struct {
	name string
	list []LanguageLabelInstruction
}

func createLanguageLabelDeclarationBuilder() LanguageLabelDeclarationBuilder {
	out := languageLabelDeclarationBuilder{
		name: "",
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageLabelDeclarationBuilder) Create() LanguageLabelDeclarationBuilder {
	return createLanguageLabelDeclarationBuilder()
}

// WithName adds a name to the builder
func (app *languageLabelDeclarationBuilder) WithName(name string) LanguageLabelDeclarationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *languageLabelDeclarationBuilder) WithInstructions(instructions []LanguageLabelInstruction) LanguageLabelDeclarationBuilder {
	app.list = instructions
	return app
}

// Now builds a new LanguageLabelDeclaration instance
func (app *languageLabelDeclarationBuilder) Now() (LanguageLabelDeclaration, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a LanguageLabelDeclaration instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 LanguageLabelInstruction instance in order to build a LanguageLabelDeclaration instance")
	}

	return createLanguageLabelDeclaration(app.name, app.list), nil
}
