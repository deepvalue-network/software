package test

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions"
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
func (app *builder) WithInstructions(ins instructions.Instructions) Builder {
	app.ins = ins
	return app
}

// Now builds a new Test instance
func (app *builder) Now() (Test, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Test instance")
	}

	if app.ins == nil {
		return nil, errors.New("the Instructions is mandatory in order to build a Test instance")
	}

	return createTest(app.name, app.ins), nil
}
