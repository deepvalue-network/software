package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

type instruction struct {
	ins     standard_instruction.Instruction
	command commands.Command
	match   match.Match
}

func createInstructionWithInstruction(
	ins standard_instruction.Instruction,
) Instruction {
	return createInstructionInternally(ins, nil, nil)
}

func createInstructionWithCommand(
	command commands.Command,
) Instruction {
	return createInstructionInternally(nil, command, nil)
}

func createInstructionWithMatch(
	match match.Match,
) Instruction {
	return createInstructionInternally(nil, nil, match)
}

func createInstructionInternally(
	ins standard_instruction.Instruction,
	command commands.Command,
	match match.Match,
) Instruction {
	out := instruction{
		ins:     ins,
		command: command,
		match:   match,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *instruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *instruction) Instruction() standard_instruction.Instruction {
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

// IsMatch returns true if there is a match, false otherwise
func (obj *instruction) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *instruction) Match() match.Match {
	return obj.match
}
