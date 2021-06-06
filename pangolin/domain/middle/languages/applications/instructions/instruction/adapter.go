package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter standard_instruction.Adapter
	commandAdapter     commands.Adapter
	matchAdapter       match.Adapter
	commonInsBuilder   CommonInstructionBuilder
	builder            Builder
}

func createAdapter(
	instructionAdapter standard_instruction.Adapter,
	commandAdapter commands.Adapter,
	matchAdapter match.Adapter,
	commonInsBuilder CommonInstructionBuilder,
	builder Builder,
) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		commandAdapter:     commandAdapter,
		matchAdapter:       matchAdapter,
		commonInsBuilder:   commonInsBuilder,
		builder:            builder,
	}

	return &out
}

// ToInstruction converts a parsed language instruction to instruction instance
func (app *adapter) ToInstruction(parsed parsers.LanguageInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if parsed.IsInstruction() {
		parsedCommonIns := parsed.Instruction()
		commonIns, err := app.ToCommonInstruction(parsedCommonIns)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(commonIns)
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

// ToCommonInstruction converts a parsed common instruction to common instruction
func (app *adapter) ToCommonInstruction(parsedCommonIns parsers.LanguageInstructionCommon) (CommonInstruction, error) {
	builder := app.commonInsBuilder.Create()
	if parsedCommonIns.IsInstruction() {
		parsedIns := parsedCommonIns.Instruction()
		ins, err := app.instructionAdapter.ToInstruction(parsedIns)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(ins)
	}

	if parsedCommonIns.IsMatch() {
		parsedMatch := parsedCommonIns.Match()
		match, err := app.matchAdapter.ToMatch(parsedMatch)
		if err != nil {
			return nil, err
		}

		builder.WithMatch(match)
	}

	return builder.Now()
}
