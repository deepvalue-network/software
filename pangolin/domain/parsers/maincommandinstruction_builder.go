package parsers

import "errors"

type mainCommandInstructionBuilder struct {
	ins    LanguageInstruction
	scopes Scopes
}

func createMainCommandInstructionBuilder() MainCommandInstructionBuilder {
	out := mainCommandInstructionBuilder{
		ins:    nil,
		scopes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *mainCommandInstructionBuilder) Create() MainCommandInstructionBuilder {
	return createMainCommandInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *mainCommandInstructionBuilder) WithInstruction(ins LanguageInstruction) MainCommandInstructionBuilder {
	app.ins = ins
	return app
}

// WithScopes add scopes to the builder
func (app *mainCommandInstructionBuilder) WithScopes(scopes Scopes) MainCommandInstructionBuilder {
	app.scopes = scopes
	return app
}

// WithScopes add scopes to the builder
func (app *mainCommandInstructionBuilder) Now() (MainCommandInstruction, error) {
	if app.ins == nil {
		return nil, errors.New("the LanguageInstruction is mandatory in order to build a MainCommandInstruction instance")
	}

	if app.scopes != nil {
		return createMainCommandInstructionWithScopes(app.ins, app.scopes), nil
	}

	return createMainCommandInstruction(app.ins), nil
}
