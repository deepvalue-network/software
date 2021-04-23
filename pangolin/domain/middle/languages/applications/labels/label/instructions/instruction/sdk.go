package instruction

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction/token"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithLabel(label label_instruction.Instruction) Builder
	WithLanguage(lang language_instruction.Instruction) Builder
	WithToken(token token.Token) Builder
	Now() (Instruction, error)
}

// Instruction represents a language label instruction
type Instruction interface {
	IsLabel() bool
	Label() label_instruction.Instruction
	IsLanguage() bool
	Language() language_instruction.Instruction
	IsToken() bool
	Token() token.Token
}
