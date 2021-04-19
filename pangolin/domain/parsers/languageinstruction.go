package parsers

type languageInstruction struct {
	ins     Instruction
	match   Match
	command Command
}

func createLanguageInstructionWithInstruction(
	ins Instruction,
) LanguageInstruction {
	return createLanguageInstructionInternally(ins, nil, nil)
}

func createLanguageInstructionWithMatch(
	match Match,
) LanguageInstruction {
	return createLanguageInstructionInternally(nil, match, nil)
}

func createLanguageInstructionWithCommand(
	command Command,
) LanguageInstruction {
	return createLanguageInstructionInternally(nil, nil, command)
}

func createLanguageInstructionInternally(
	ins Instruction,
	match Match,
	command Command,
) LanguageInstruction {
	out := languageInstruction{
		ins:     ins,
		match:   match,
		command: command,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *languageInstruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *languageInstruction) Instruction() Instruction {
	return obj.ins
}

// IsMatch returns true if there is a match, false otherwise
func (obj *languageInstruction) IsMatch() bool {
	return obj.match != nil
}

// Match returns a match, if any
func (obj *languageInstruction) Match() Match {
	return obj.match
}

// IsCommand returns true if there is a command, false otherwise
func (obj *languageInstruction) IsCommand() bool {
	return obj.command != nil
}

// Command returns a command, if any
func (obj *languageInstruction) Command() Command {
	return obj.command
}
