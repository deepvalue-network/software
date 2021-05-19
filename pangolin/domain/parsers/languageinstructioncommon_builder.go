package parsers

import "errors"

type languageInstructionCommonBuilder struct {
	ins   Instruction
	match Match
}

func createLanguageInstructionCommonBuilder() LanguageInstructionCommonBuilder {
	out := languageInstructionCommonBuilder{
		ins:   nil,
		match: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageInstructionCommonBuilder) Create() LanguageInstructionCommonBuilder {
	return createLanguageInstructionCommonBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *languageInstructionCommonBuilder) WithInstruction(ins Instruction) LanguageInstructionCommonBuilder {
	app.ins = ins
	return app
}

// WithMatch adds a match to the builder
func (app *languageInstructionCommonBuilder) WithMatch(match Match) LanguageInstructionCommonBuilder {
	app.match = match
	return app
}

// Now builds a new LanguageInstructionCommon instance
func (app *languageInstructionCommonBuilder) Now() (LanguageInstructionCommon, error) {
	if app.ins != nil {
		return createLanguageInstructionCommonWithInstruction(app.ins), nil
	}

	if app.match != nil {
		return createLanguageInstructionCommonWithMatch(app.match), nil
	}

	return nil, errors.New("the LanguageInstructionCommon is invalid")
}
