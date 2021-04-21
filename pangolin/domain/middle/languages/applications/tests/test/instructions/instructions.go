package instructions

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"

type instructions struct {
	list []instruction.Instruction
}

func createInstructions(
	list []instruction.Instruction,
) Instructions {
	out := instructions{
		list: list,
	}

	return &out
}

// All returns the instructions
func (obj *instructions) All() []instruction.Instruction {
	return obj.list
}
