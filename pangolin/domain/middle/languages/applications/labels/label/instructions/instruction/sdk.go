package instruction

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction/token"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	labelAdapter := label_instruction.NewAdapter()
	languageAdapter := language_instruction.NewAdapter()
	tokenAdapter := token.NewAdapter()
	builder := NewBuilder()
	return createAdapter(labelAdapter, languageAdapter, tokenAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an instruction adapter
type Adapter interface {
	ToInstruction(parsed parsers.LanguageLabelInstruction) (Instruction, error)
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
