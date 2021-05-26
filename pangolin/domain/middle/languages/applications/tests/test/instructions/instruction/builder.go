package instruction

import (
	"errors"

	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
)

type builder struct {
	lang          standard_instruction.CommonInstruction
	test          test_instruction.Instruction
	isInterpret bool
}

func createBuilder() Builder {
	out := builder{
		lang:          nil,
		test:          nil,
		isInterpret: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(lang standard_instruction.CommonInstruction) Builder {
	app.lang = lang
	return app
}

// WithTest adds a test to the builder
func (app *builder) WithTest(test test_instruction.Instruction) Builder {
	app.test = test
	return app
}

// IsInterpret flags the builder as interpret
func (app *builder) IsInterpret() Builder {
	app.isInterpret = true
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.lang != nil {
		return createInstructionWithLanguage(app.lang), nil
	}

	if app.test != nil {
		return createInstructionWithTest(app.test), nil
	}

	if app.isInterpret {
		return createInstructionWithInterpret(), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
