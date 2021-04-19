package instructions

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"

// Instructions represents instructions
type Instructions interface {
	Instructions() []instruction.Instruction
}
