package instruction

import (
	"errors"

	ins "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
)

type builder struct {
	isRet bool
	ins   ins.Instruction
}

func createBuilder() Builder {
	out := builder{
		isRet: false,
		ins:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsRet flags the builder as return
func (app *builder) IsRet() Builder {
	app.isRet = true
	return app
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins ins.Instruction) Builder {
	app.ins = ins
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.isRet {
		return createInstructionWithReturn(), nil
	}

	if app.ins != nil {
		return createInstructionWithInstruction(app.ins), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
