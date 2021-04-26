package mains

import (
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
)

type instruction struct {
	ins    application_instruction.Instruction
	scopes []bool
}

func createInstruction(
	ins application_instruction.Instruction,
) Instruction {
	return createInstructionInternally(ins, nil)
}

func createInstructionWithScopes(
	ins application_instruction.Instruction,
	scopes []bool,
) Instruction {
	return createInstructionInternally(ins, scopes)
}

func createInstructionInternally(
	ins application_instruction.Instruction,
	scopes []bool,
) Instruction {
	out := instruction{
		ins:    ins,
		scopes: scopes,
	}

	return &out
}

// Instruction returns the instruction
func (obj *instruction) Instruction() application_instruction.Instruction {
	return obj.ins
}

// HasScopes returns true if there is scopes, false otherwise
func (obj *instruction) HasScopes() bool {
	return obj.scopes != nil
}

// Scopes returns the scopes, if any (external = true, internal = false)
func (obj *instruction) Scopes() []bool {
	return obj.scopes
}
