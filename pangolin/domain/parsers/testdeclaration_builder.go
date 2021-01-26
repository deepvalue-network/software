package parsers

import "errors"

type testDeclarationBuilder struct {
	name         string
	instructions []TestInstruction
}

func createTestDeclarationBuilder() TestDeclarationBuilder {
	out := testDeclarationBuilder{
		name:         "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testDeclarationBuilder) Create() TestDeclarationBuilder {
	return createTestDeclarationBuilder()
}

// WithName adds a name to the builder
func (app *testDeclarationBuilder) WithName(name string) TestDeclarationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *testDeclarationBuilder) WithInstructions(instructions []TestInstruction) TestDeclarationBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new TestDeclaration instance
func (app *testDeclarationBuilder) Now() (TestDeclaration, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a TestDeclaration instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the []TestInstruction are mandatory in order to build a TestDeclaration instance")
	}

	return createTestDeclaration(app.name, app.instructions), nil
}
