package label

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"
)

type builder struct {
	name string
	ins  instructions.Instructions
}

func createBuilder() Builder {
	out := builder{
		name: "",
		ins:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.ins = instructions
	return app
}

// Now builds a new Label instance
func (app *builder) Now() (Label, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Label instance")
	}

	if app.ins == nil {
		return nil, errors.New("the instructions are mandatory in order to build a Label instance")
	}

	return createLabel(app.name, app.ins), nil
}
