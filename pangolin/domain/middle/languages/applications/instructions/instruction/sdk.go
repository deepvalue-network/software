package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
)

// Instruction represents a language application instruction
type Instruction interface {
	IsInstruction() bool
	Instruction() instruction.Instruction
	IsMatch() bool
	Match() match.Match
}
