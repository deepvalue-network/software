package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter standard_instruction.Adapter
	commandAdapter     commands.Adapter
	matchAdapter       match.Adapter
	builder            Builder
}

func createAdapter(
	instructionAdapter standard_instruction.Adapter,
	commandAdapter commands.Adapter,
	matchAdapter match.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		commandAdapter:     commandAdapter,
		matchAdapter:       matchAdapter,
		builder:            builder,
	}

	return &out
}

// ToInstruction converts a parsed language instruction to instruction instance
func (app *adapter) ToInstruction(parsed parsers.LanguageInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if parsed.IsInstruction() {
		parsedIns := parsed.Instruction()
		ins, err := app.instructionAdapter.ToInstruction(parsedIns)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(ins)
	}

	if parsed.IsMatch() {
		parsedMatch := parsed.Match()
		match, err := app.matchAdapter.ToMatch(parsedMatch)
		if err != nil {
			return nil, err
		}

		builder.WithMatch(match)
	}

	if parsed.IsCommand() {
		parsedCommand := parsed.Command()
		command, err := app.commandAdapter.ToCommand(parsedCommand)
		if err != nil {
			return nil, err
		}

		builder.WithCommand(command)
	}

	return builder.Now()
}
