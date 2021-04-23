package label

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionsAdapter instructions.Adapter
	builder             Builder
}

func createAdapter(instructionsAdapter instructions.Adapter, builder Builder) Adapter {
	out := adapter{
		instructionsAdapter: instructionsAdapter,
		builder:             builder,
	}

	return &out
}

// ToLabel converts a parsed label declaration to an optmized Label instance
func (app *adapter) ToLabel(declaration parsers.LabelDeclaration) (Label, error) {
	name := declaration.Name()
	parsedInstructions := declaration.Instructions()
	instructions, err := app.instructionsAdapter.ToInstructions(parsedInstructions)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithName(name).WithInstructions(instructions).Now()
}
