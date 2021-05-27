package instructions

import "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"

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
