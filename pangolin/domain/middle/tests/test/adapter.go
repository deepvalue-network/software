package test

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/tests/test/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionsAdapter instructions.Adapter
	builder             Builder
}

func createAdapter(
	instructionsAdapter instructions.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		instructionsAdapter: instructionsAdapter,
		builder:             builder,
	}

	return &out
}

// ToTest converts a TestDeclaration to a Test instance
func (app *adapter) ToTest(declaration parsers.TestDeclaration) (Test, error) {
	name := declaration.Name()
	parsedIns := declaration.Instructions()
	ins, err := app.instructionsAdapter.ToInstructions(parsedIns)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithName(name).WithInstructions(ins).Now()
}
