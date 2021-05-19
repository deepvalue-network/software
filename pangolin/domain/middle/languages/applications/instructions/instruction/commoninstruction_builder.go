package instruction

import (
	"errors"

	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

type commonInstructionBuilder struct {
	ins   standard_instruction.Instruction
	match match.Match
}

func createCommonInstructionBuilder() CommonInstructionBuilder {
	out := commonInstructionBuilder{
		ins:   nil,
		match: nil,
	}

	return &out
}

// Create initializes the builder
func (app *commonInstructionBuilder) Create() CommonInstructionBuilder {
	return createCommonInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *commonInstructionBuilder) WithInstruction(ins standard_instruction.Instruction) CommonInstructionBuilder {
	app.ins = ins
	return app
}

// WithMatch adds a match to the builder
func (app *commonInstructionBuilder) WithMatch(match match.Match) CommonInstructionBuilder {
	app.match = match
	return app
}

// Now builds a new Instruction instance
func (app *commonInstructionBuilder) Now() (CommonInstruction, error) {
	if app.ins != nil {
		return createCommonInstructionWithInstruction(app.ins), nil
	}

	if app.match != nil {
		return createCommonInstructionWithMatch(app.match), nil
	}

	return nil, errors.New("the CommonInstruction is invalid")
}
