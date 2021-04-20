package mains

import (
	"errors"

	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
)

type instructionBuilder struct {
	ins    language_instruction.Instruction
	scopes []bool
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		ins:    nil,
		scopes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *instructionBuilder) WithInstruction(ins language_instruction.Instruction) InstructionBuilder {
	app.ins = ins
	return app
}

// WithScopes add scopes to the builder
func (app *instructionBuilder) WithScopes(scopes []bool) InstructionBuilder {
	app.scopes = scopes
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.ins == nil {
		return nil, errors.New("the label instruction is mandatory in order to build an Instruction instance")
	}

	if app.scopes != nil {
		return createInstructionWithScopes(app.ins, app.scopes), nil
	}

	return createInstruction(app.ins), nil
}
