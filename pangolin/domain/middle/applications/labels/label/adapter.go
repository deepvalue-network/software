package label

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions"
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

// ToLabel converts a parsed language label declaration to label instance
func (app *adapter) ToLabel(parsed parsers.LanguageLabelDeclaration) (Label, error) {
	parsedInstructions := parsed.Instructions()
	instructions, err := app.instructionsAdapter.ToInstructions(parsedInstructions)
	if err != nil {
		return nil, err
	}

	name := parsed.Name()
	return app.builder.Create().WithName(name).WithInstructions(instructions).Now()
}
