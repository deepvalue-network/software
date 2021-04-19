package instructions

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"

// Instructions represents instructions
type Instructions interface {
	All() []instruction.Instruction
}
