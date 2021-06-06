package instruction

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction/token"
)

type instruction struct {
	label label_instruction.Instruction
	lang  language_instruction.Instruction
	tok   token.Token
}

func createInstructionWithLabel(
	label label_instruction.Instruction,
) Instruction {
	return createInstructionInternally(label, nil, nil)
}

func createInstructionWithLanguage(
	lang language_instruction.Instruction,
) Instruction {
	return createInstructionInternally(nil, lang, nil)
}

func createInstructionWithToken(
	tok token.Token,
) Instruction {
	return createInstructionInternally(nil, nil, tok)
}

func createInstructionInternally(
	label label_instruction.Instruction,
	lang language_instruction.Instruction,
	tok token.Token,
) Instruction {
	out := instruction{
		label: label,
		lang:  lang,
		tok:   tok,
	}

	return &out
}

// IsLabel returns true if there is a label, false otherwise
func (obj *instruction) IsLabel() bool {
	return obj.label != nil
}

// Label returns the label, if any
func (obj *instruction) Label() label_instruction.Instruction {
	return obj.label
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *instruction) IsLanguage() bool {
	return obj.lang != nil
}

// Language returns the language, if any
func (obj *instruction) Language() language_instruction.Instruction {
	return obj.lang
}

// IsToken returns true if there is a token, false otherwise
func (obj *instruction) IsToken() bool {
	return obj.tok != nil
}

// Token returns the token, if any
func (obj *instruction) Token() token.Token {
	return obj.tok
}
