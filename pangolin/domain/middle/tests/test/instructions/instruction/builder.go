package instruction

import (
	"errors"

	ins "github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"
)

type builder struct {
	isStart bool
	isStop  bool
	ins     ins.Instruction
}

func createBuilder() Builder {
	out := builder{
		isStart: false,
		isStop:  false,
		ins:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsStart sets the builder as start
func (app *builder) IsStart() Builder {
	app.isStart = true
	return app
}

// IsStop sets the builder as stop
func (app *builder) IsStop() Builder {
	app.isStop = true
	return app
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins ins.Instruction) Builder {
	app.ins = ins
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.isStart {
		return createInstructionWithStart(), nil
	}

	if app.isStop {
		return createInstructionWithStop(), nil
	}

	if app.ins != nil {
		return createInstructionWithInstruction(app.ins), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
