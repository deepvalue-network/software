package instruction

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction/token"
)

// Instruction represents a language label instruction
type Instruction interface {
	IsLabel() bool
	Label() label_instruction.Instruction
	IsLanguage() bool
	Language() language_instruction.Instruction
	IsToken() bool
	Token() token.Token
}
