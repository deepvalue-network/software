package instruction

import (
	"errors"

	ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
)

type builder struct {
	isAssert bool
	readFile ReadFile
	ins      ins.Instruction
}

func createBuilder() Builder {
	out := builder{
		isAssert: false,
		readFile: nil,
		ins:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsAssert sets the builder as assert
func (app *builder) IsAssert() Builder {
	app.isAssert = true
	return app
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins ins.Instruction) Builder {
	app.ins = ins
	return app
}

// WithReadFile adds a readFile to the builder
func (app *builder) WithReadFile(readFile ReadFile) Builder {
	app.readFile = readFile
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.isAssert {
		return createInstructionWithAssert(), nil
	}

	if app.readFile != nil {
		return createInstructionWithReadFile(app.readFile), nil
	}

	if app.ins != nil {
		return createInstructionWithInstruction(app.ins), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
