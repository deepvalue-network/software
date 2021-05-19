package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

type commonInstruction struct {
	ins   standard_instruction.Instruction
	match match.Match
}

func createCommonInstructionWithInstruction(
	ins standard_instruction.Instruction,
) CommonInstruction {
	return createCommonInstructionInternally(ins, nil)
}

func createCommonInstructionWithMatch(
	match match.Match,
) CommonInstruction {
	return createCommonInstructionInternally(nil, match)
}

func createCommonInstructionInternally(
	ins standard_instruction.Instruction,
	match match.Match,
) CommonInstruction {
	out := commonInstruction{
		ins:   ins,
		match: match,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *commonInstruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction retruns the instruction, if any
func (obj *commonInstruction) Instruction() standard_instruction.Instruction {
	return obj.ins
}

// IsMatch returns true if there is a match, false otherwise
func (obj *commonInstruction) IsMatch() bool {
	return obj.match != nil
}

// Match retruns the match, if any
func (obj *commonInstruction) Match() match.Match {
	return obj.match
}
