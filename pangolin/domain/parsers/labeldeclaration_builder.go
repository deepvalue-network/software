package parsers

import "errors"

type labelDeclarationBuilder struct {
	name         string
	instructions []LabelInstruction
}

func createLabelDeclarationBuilder() LabelDeclarationBuilder {
	out := labelDeclarationBuilder{
		name:         "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *labelDeclarationBuilder) Create() LabelDeclarationBuilder {
	return createLabelDeclarationBuilder()
}

// WithName adds a name to the builder
func (app *labelDeclarationBuilder) WithName(name string) LabelDeclarationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *labelDeclarationBuilder) WithInstructions(ins []LabelInstruction) LabelDeclarationBuilder {
	app.instructions = ins
	return app
}

// Now builds a new LabelDeclaration instance
func (app *labelDeclarationBuilder) Now() (LabelDeclaration, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a LabelDeclaration instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the []Instruction are mandatory in order to build a LabelDeclaration instance")
	}

	return createLabelDeclaration(app.name, app.instructions), nil
}
