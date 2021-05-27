package instruction

import (
	"errors"

	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction/token"
)

type builder struct {
	label label_instruction.Instruction
	lang  language_instruction.Instruction
	tok   token.Token
}

func createBuilder() Builder {
	out := builder{
		label: nil,
		lang:  nil,
		tok:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLabel adds a label to the builder
func (app *builder) WithLabel(label label_instruction.Instruction) Builder {
	app.label = label
	return app
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(lang language_instruction.Instruction) Builder {
	app.lang = lang
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token token.Token) Builder {
	app.tok = token
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.label != nil {
		return createInstructionWithLabel(app.label), nil
	}

	if app.lang != nil {
		return createInstructionWithLanguage(app.lang), nil
	}

	if app.tok != nil {
		return createInstructionWithToken(app.tok), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
