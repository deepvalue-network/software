package parsers

import "errors"

type labelCommandInstructionBuilder struct {
	ins    LabelInstruction
	scopes Scopes
}

func createLabelCommandInstructionBuilder() LabelCommandInstructionBuilder {
	out := labelCommandInstructionBuilder{
		ins:    nil,
		scopes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *labelCommandInstructionBuilder) Create() LabelCommandInstructionBuilder {
	return createLabelCommandInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *labelCommandInstructionBuilder) WithInstruction(ins LabelInstruction) LabelCommandInstructionBuilder {
	app.ins = ins
	return app
}

// WithScopes add scopes to the builder
func (app *labelCommandInstructionBuilder) WithScopes(scopes Scopes) LabelCommandInstructionBuilder {
	app.scopes = scopes
	return app
}

// Now builds a new LabelCommandInstruction instance
func (app *labelCommandInstructionBuilder) Now() (LabelCommandInstruction, error) {
	if app.ins == nil {
		return nil, errors.New("the LabelInstruction is mandatory in order to build a LabelCommandInstruction instance")
	}

	if app.scopes != nil {
		return createLabelCommandInstructionWithScopes(app.ins, app.scopes), nil
	}

	return createLabelCommandInstruction(app.ins), nil
}
