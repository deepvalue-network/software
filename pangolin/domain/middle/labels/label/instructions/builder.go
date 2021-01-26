package instructions

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/middle/labels/label/instructions/instruction"
)

type builder struct {
	lst []instruction.Instruction
	mp  map[string]instruction.Instruction
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
		mp:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add list to the builder
func (app *builder) WithList(lst []instruction.Instruction) Builder {
	app.lst = lst
	return app
}

// WithMap add map to the builder
func (app *builder) WithMap(mp map[string]instruction.Instruction) Builder {
	app.mp = mp
	return app
}

// Now builds a new Instructions instance
func (app *builder) Now() (Instructions, error) {
	if app.mp != nil {
		lst := []instruction.Instruction{}
		for _, oneInstruction := range app.mp {
			lst = append(lst, oneInstruction)
		}

		app.lst = lst
	}

	if app.lst == nil {
		return nil, errors.New("the list is mandatory in order to build a Instructions instance")
	}

	return createInstructions(app.lst), nil
}
