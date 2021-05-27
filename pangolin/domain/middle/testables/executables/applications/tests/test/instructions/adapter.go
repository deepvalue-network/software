package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions/instruction"
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

// ToInstructions converts parsed TestInstruction to an optimized test Instructions instance
func (app *adapter) ToInstructions(testInstructions []parsers.TestInstruction) (Instructions, error) {
	lst := []instruction.Instruction{}
	for _, oneTestInstruction := range testInstructions {
		testIns, err := app.instructionAdapter.ToInstruction(oneTestInstruction)
		if err != nil {
			return nil, err
		}

		lst = append(lst, testIns)
	}

	return app.builder.Create().WithList(lst).Now()
}
