package instructions

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
)

type builder struct {
	list []instruction.Instruction
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add instructions to the builder
func (app *builder) WithList(list []instruction.Instruction) Builder {
	app.list = list
	return app
}

// Now builds a new Instructins instance
func (app *builder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the []Instruction are mandatory in order to build an Instructions instance")
	}

	return createInstructions(app.list), nil
}
