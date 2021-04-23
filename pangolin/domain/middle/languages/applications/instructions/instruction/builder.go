package instruction

import (
	"errors"

	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

type builder struct {
	ins   standard_instruction.Instruction
	match match.Match
}

func createBuilder() Builder {
	out := builder{
		ins:   nil,
		match: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins standard_instruction.Instruction) Builder {
	app.ins = ins
	return app
}

// WithMatch adds a match to the builder
func (app *builder) WithMatch(match match.Match) Builder {
	app.match = match
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.ins != nil {
		return createInstructionWithInstruction(app.ins), nil
	}

	if app.match != nil {
		return createInstructionWithMatch(app.match), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
