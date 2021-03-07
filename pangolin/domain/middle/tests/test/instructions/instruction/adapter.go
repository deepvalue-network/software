package instruction

import (
	ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	builder            Builder
	readFileBuilder    ReadFileBuilder
	instructionAdapter ins.Adapter
}

func createAdapter(
	builder Builder,
	readFileBuilder ReadFileBuilder,
	instructionAdapter ins.Adapter,
) Adapter {
	out := adapter{
		builder:            builder,
		readFileBuilder:    readFileBuilder,
		instructionAdapter: instructionAdapter,
	}

	return &out
}

// ToInstruction converts a testInstruction to an Instruction instance
func (app *adapter) ToInstruction(testInstruction parsers.TestInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if testInstruction.IsInstruction() {
		parsedIns := testInstruction.Instruction()
		ins, err := app.instructionAdapter.ToInstruction(parsedIns)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(ins)
	}

	if testInstruction.IsReadFile() {
		parsedReadFile := testInstruction.ReadFile()
		variable := parsedReadFile.Variable().String()
		path := parsedReadFile.Path().String()
		ins, err := app.readFileBuilder.Create().WithVariable(variable).WithPath(path).Now()
		if err != nil {
			return nil, err
		}

		builder.WithReadFile(ins)
	}

	if testInstruction.IsAssert() {
		builder.IsAssert()
	}

	return builder.Now()
}
