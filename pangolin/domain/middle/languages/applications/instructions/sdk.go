package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
)

// Instructions represents a language instructions
type Instructions interface {
	All() []instruction.Instruction
}
