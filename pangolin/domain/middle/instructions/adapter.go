package instructions

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter instruction.Adapter
	builder            Builder
}

func createAdapter(instructionAdapter instruction.Adapter, builder Builder) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		builder:            builder,
	}

	return &out
}

// ToInstructions converts parsed instructions to an optimized Instructions instance
func (app *adapter) ToInstructions(parsed []parsers.Instruction) (Instructions, error) {
	lst := []instruction.Instruction{}
	for _, oneParsedInstruction := range parsed {
		ins, err := app.instructionAdapter.ToInstruction(oneParsedInstruction)
		if err != nil {
			return nil, err
		}

		lst = append(lst, ins)
	}

	return app.builder.Create().WithList(lst).Now()
}
