package parsers

import "errors"

type languageInstructionBuilder struct {
	ins     Instruction
	match   Match
	command Command
}

func createLanguageInstructionBuilder() LanguageInstructionBuilder {
	out := languageInstructionBuilder{
		ins:     nil,
		match:   nil,
		command: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageInstructionBuilder) Create() LanguageInstructionBuilder {
	return createLanguageInstructionBuilder()
}

// WithInstruction adds an instruction to the builder
func (app *languageInstructionBuilder) WithInstruction(ins Instruction) LanguageInstructionBuilder {
	app.ins = ins
	return app
}

// WithMatch adds a match to the builder
func (app *languageInstructionBuilder) WithMatch(match Match) LanguageInstructionBuilder {
	app.match = match
	return app
}

// WithCommand adds a command to the builder
func (app *languageInstructionBuilder) WithCommand(command Command) LanguageInstructionBuilder {
	app.command = command
	return app
}

// Now builds a new LanguageInstruction instance
func (app *languageInstructionBuilder) Now() (LanguageInstruction, error) {
	if app.ins != nil {
		return createLanguageInstructionWithInstruction(app.ins), nil
	}

	if app.match != nil {
		return createLanguageInstructionWithMatch(app.match), nil
	}

	if app.command != nil {
		return createLanguageInstructionWithCommand(app.command), nil
	}

	return nil, errors.New("the LanguageInstruction instance is invalid")
}
