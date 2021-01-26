package parsers

import "errors"

type labelInstructionBuilder struct {
	isRet       bool
	instruction Instruction
}

func createLabelInstructionBuilder() LabelInstructionBuilder {
	out := labelInstructionBuilder{
		isRet:       false,
		instruction: nil,
	}

	return &out
}

// Create initializes the builder
func (app *labelInstructionBuilder) Create() LabelInstructionBuilder {
	return createLabelInstructionBuilder()
}

// IsRet flags the builder as return
func (app *labelInstructionBuilder) IsRet() LabelInstructionBuilder {
	app.isRet = true
	return app
}

// WithInstruction adds an instruction to the builder
func (app *labelInstructionBuilder) WithInstruction(ins Instruction) LabelInstructionBuilder {
	app.instruction = ins
	return app
}

// Now builds a new LabelInstruction instance
func (app *labelInstructionBuilder) Now() (LabelInstruction, error) {
	if app.isRet {
		return createLabelInstructionWithRet(), nil
	}

	if app.instruction != nil {
		return createLabelInstructionWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the LabelInstruction is invalid")
}
