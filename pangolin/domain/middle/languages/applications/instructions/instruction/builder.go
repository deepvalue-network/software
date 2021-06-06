package instruction

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
)

type builder struct {
	ins     CommonInstruction
	command commands.Command
}

func createBuilder() Builder {
	out := builder{
		ins:     nil,
		command: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins CommonInstruction) Builder {
	app.ins = ins
	return app
}

// WithCommand adds a command to the builder
func (app *builder) WithCommand(command commands.Command) Builder {
	app.command = command
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.ins != nil {
		return createInstructionWithInstruction(app.ins), nil
	}

	if app.command != nil {
		return createInstructionWithCommand(app.command), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
