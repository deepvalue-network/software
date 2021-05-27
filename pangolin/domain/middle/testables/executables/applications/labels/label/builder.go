package label

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions"
)

type builder struct {
	instructionsBuilder instructions.Builder
	name                string
	ins                 instructions.Instructions
}

func createBuilder(instructionsBuilder instructions.Builder) Builder {
	out := builder{
		instructionsBuilder: instructionsBuilder,
		name:                "",
		ins:                 nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.instructionsBuilder)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithName adds a name to the builder
func (app *builder) WithInstructions(ins instructions.Instructions) Builder {
	app.ins = ins
	return app
}

// Now builds a new Label instance
func (app *builder) Now() (Label, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Label instance")
	}

	if app.ins == nil {
		ins, err := app.instructionsBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.ins = ins
	}

	return createLabel(app.name, app.ins), nil
}
