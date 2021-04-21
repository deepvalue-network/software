package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

type instruction struct {
	ins   standard_instruction.Instruction
	match match.Match
}

func createInstructionWithInstruction(
	ins standard_instruction.Instruction,
) Instruction {
	return createInstructionInternally(ins, nil)
}

func createInstructionWithMatch(
	match match.Match,
) Instruction {
	return createInstructionInternally(nil, match)
}

func createInstructionInternally(
	ins standard_instruction.Instruction,
	match match.Match,
) Instruction {
	out := instruction{
		ins:   ins,
		match: match,
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

// IsMatch returns true if there is a match, false otherwise
func (obj *instruction) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *instruction) Match() match.Match {
	return obj.match
}
