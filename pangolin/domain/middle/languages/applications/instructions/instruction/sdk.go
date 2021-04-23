package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents instructions adapter
type Adapter interface {
	ToInstruction(parsed parsers.LanguageInstruction) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithInstruction(ins standard_instruction.Instruction) Builder
	WithMatch(match match.Match) Builder
	Now() (Instruction, error)
}

// Instruction represents a language application instruction
type Instruction interface {
	IsInstruction() bool
	Instruction() standard_instruction.Instruction
	IsMatch() bool
	Match() match.Match
}
