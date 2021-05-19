package parsers

type languageInstruction struct {
	ins     LanguageInstructionCommon
	command Command
}

func createLanguageInstructionWithInstruction(
	ins LanguageInstructionCommon,
) LanguageInstruction {
	return createLanguageInstructionInternally(ins, nil)
}

func createLanguageInstructionWithCommand(
	command Command,
) LanguageInstruction {
	return createLanguageInstructionInternally(nil, command)
}

func createLanguageInstructionInternally(
	ins LanguageInstructionCommon,
	command Command,
) LanguageInstruction {
	out := languageInstruction{
		ins:     ins,
		command: command,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *languageInstruction) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *languageInstruction) Instruction() LanguageInstructionCommon {
	return obj.ins
}

// IsCommand returns true if there is a command, false otherwise
func (obj *languageInstruction) IsCommand() bool {
	return obj.command != nil
}

// Command returns a command, if any
func (obj *languageInstruction) Command() Command {
	return obj.command
}
