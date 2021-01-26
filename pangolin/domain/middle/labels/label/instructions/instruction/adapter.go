package instruction

import (
	ins "github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter ins.Adapter
	builder            Builder
}

func createAdapter(instructionAdapter ins.Adapter, builder Builder) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		builder:            builder,
	}

	return &out
}

// ToInstruction converts a parsed instruction to an optimized instruction
func (app *adapter) ToInstruction(parsed parsers.LabelInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if parsed.IsRet() {
		builder.IsRet()
	}

	if parsed.IsInstruction() {
		parsedInstruction := parsed.Instruction()
		ins, err := app.instructionAdapter.ToInstruction(parsedInstruction)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(ins)
	}

	return builder.Now()
}
