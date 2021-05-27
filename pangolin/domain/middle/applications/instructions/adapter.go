package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
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

// ToInstructions converts a parsed language main section to instructions instance
func (app *adapter) ToInstructions(parsed parsers.LanguageMainSection) (Instructions, error) {
	instructions := []instruction.Instruction{}
	parsedLangInstructions := parsed.Instructions()
	for _, oneLangIns := range parsedLangInstructions {
		ins, err := app.instructionAdapter.ToInstruction(oneLangIns)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, ins)
	}

	return app.builder.Create().WithList(instructions).Now()
}
