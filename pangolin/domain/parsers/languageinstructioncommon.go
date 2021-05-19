package parsers

type languageInstructionCommon struct {
	ins   Instruction
	match Match
}

func createLanguageInstructionCommonWithInstruction(
	ins Instruction,
) LanguageInstructionCommon {
	return createLanguageInstructionCommonInternally(ins, nil)
}

func createLanguageInstructionCommonWithMatch(
	match Match,
) LanguageInstructionCommon {
	return createLanguageInstructionCommonInternally(nil, match)
}

func createLanguageInstructionCommonInternally(
	ins Instruction,
	match Match,
) LanguageInstructionCommon {
	out := languageInstructionCommon{
		ins:   ins,
		match: match,
	}

	return &out
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *languageInstructionCommon) IsInstruction() bool {
	return obj.ins != nil
}

// Instruction returns the instruction, if any
func (obj *languageInstructionCommon) Instruction() Instruction {
	return obj.ins
}

// IsMatch retruns true if there is a match, false otherwise
func (obj *languageInstructionCommon) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *languageInstructionCommon) Match() Match {
	return obj.match
}
