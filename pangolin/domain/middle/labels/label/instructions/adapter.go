package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
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

// ToInstructions converts parsed label instructions to an optimized label Instructions
func (app *adapter) ToInstructions(instructions []parsers.LabelInstruction) (Instructions, error) {
	lst := []instruction.Instruction{}
	for _, oneInstruction := range instructions {
		ins, err := app.instructionAdapter.ToInstruction(oneInstruction)
		if err != nil {
			return nil, err
		}

		lst = append(lst, ins)
	}

	return app.builder.Create().WithList(lst).Now()
}
