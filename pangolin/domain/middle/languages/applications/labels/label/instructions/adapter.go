package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter instruction.Adapter
	builder            Builder
}

func createAdapter(
	instructionAdapter instruction.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		builder:            builder,
	}

	return &out
}

// ToInstructions converts parsed language label instructions to instructions instance
func (app *adapter) ToInstructions(parsed []parsers.LanguageLabelInstruction) (Instructions, error) {
	out := []instruction.Instruction{}
	for _, oneParsedIns := range parsed {
		ins, err := app.instructionAdapter.ToInstruction(oneParsedIns)
		if err != nil {
			return nil, err
		}

		out = append(out, ins)
	}

	return app.builder.Create().WithList(out).Now()
}
