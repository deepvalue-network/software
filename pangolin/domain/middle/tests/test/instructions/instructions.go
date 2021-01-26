package instructions

import "github.com/steve-care-software/products/pangolin/domain/middle/tests/test/instructions/instruction"

type instructions struct {
	lst []instruction.Instruction
}

func createInstructions(lst []instruction.Instruction) Instructions {
	out := instructions{
		lst: lst,
	}

	return &out
}

// All return all the instructions
func (obj *instructions) All() []instruction.Instruction {
	return obj.lst
}
