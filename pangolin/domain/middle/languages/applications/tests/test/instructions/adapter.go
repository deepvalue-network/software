package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
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

// ToInstructions converts a language test instruction to instructions instance
func (app *adapter) ToInstructions(parsed []parsers.LanguageTestInstruction) (Instructions, error) {
	instructions := []instruction.Instruction{}
	for _, oneParsedIns := range parsed {
		ins, err := app.instructionAdapter.ToInstruction(oneParsedIns)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, ins)
	}

	return app.builder.Create().WithList(instructions).Now()
}
