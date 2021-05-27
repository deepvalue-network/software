package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands"
)

type instruction struct {
	ins     CommonInstruction
	command commands.Command
}

func createInstructionWithInstruction(
	ins CommonInstruction,
) Instruction {
	return createInstructionInternally(ins, nil)
}

func createInstructionWithCommand(
	command commands.Command,
) Instruction {
	return createInstructionInternally(nil, command)
}

func createInstructionInternally(
	ins CommonInstruction,
	command commands.Command,
) Instruction {
	out := instruction{
		ins:     ins,
		command: command,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *instruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *instruction) Instruction() CommonInstruction {
	return obj.ins
}

// IsCommand returns true if there is a command, false otherwise
func (obj *instruction) IsCommand() bool {
	return obj.command != nil
}

// Command returns the command, if any
func (obj *instruction) Command() commands.Command {
	return obj.command
}
