package parsers

type labelCommandInstruction struct {
	ins    LanguageLabelInstruction
	scopes Scopes
}

func createLabelCommandInstruction(
	ins LanguageLabelInstruction,
) LabelCommandInstruction {
	return createLabelCommandInstructionInternally(ins, nil)
}

func createLabelCommandInstructionWithScopes(
	ins LanguageLabelInstruction,
    scopes Scopes,
) LabelCommandInstruction {
	return createLabelCommandInstructionInternally(ins, scopes)
}

func createLabelCommandInstructionInternally(
	ins LanguageLabelInstruction,
	scopes Scopes,
) LabelCommandInstruction {
	out := labelCommandInstruction{
		ins:    ins,
		scopes: scopes,
	}

	return &out
}

// Instruction returns the instruction
func (obj *labelCommandInstruction) Instruction() LanguageLabelInstruction {
	return obj.ins
}

// HasScopes returns true if there is scopes, false otherwise
func (obj *labelCommandInstruction) HasScopes() bool {
	return obj.scopes != nil
}

// Scopes returns the scopes, if any
func (obj *labelCommandInstruction) Scopes() Scopes {
	return obj.scopes
}
