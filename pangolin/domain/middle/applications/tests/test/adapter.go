package test

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions"
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

// ToTest converts a language test declaration to a Test instance
func (app *adapter) ToTest(parsed parsers.LanguageTestDeclaration) (Test, error) {
	parsedInstructions := parsed.Instructions()
	instructions, err := app.instructionsAdapter.ToInstructions(parsedInstructions)
	if err != nil {
		return nil, err
	}

	name := parsed.Name()
	return app.builder.Create().WithName(name).WithInstructions(instructions).Now()
}
