package instructions

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents instructions adapter
type Adapter interface {
	ToInstructions(parsed parsers.LanguageMainSection) (Instructions, error)
}

// Builder represents an instructions builder
type Builder interface {
	Create() Builder
	WithList(list []instruction.Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents a language instructions
type Instructions interface {
	All() []instruction.Instruction
}
