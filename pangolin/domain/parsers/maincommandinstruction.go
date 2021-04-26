package parsers

type mainCommandInstruction struct {
	ins    Instruction
	scopes Scopes
}

func createMainCommandInstruction(
	ins Instruction,
) MainCommandInstruction {
	return createMainCommandInstructionInternally(ins, nil)
}

func createMainCommandInstructionWithScopes(
	ins Instruction,
	scopes Scopes,
) MainCommandInstruction {
	return createMainCommandInstructionInternally(ins, scopes)
}

func createMainCommandInstructionInternally(
	ins Instruction,
	scopes Scopes,
) MainCommandInstruction {
	out := mainCommandInstruction{
		ins:    ins,
		scopes: scopes,
	}

	return &out
}

// Instruction returns the instruction
func (obj *mainCommandInstruction) Instruction() Instruction {
	return obj.ins
}

// HasScopes returns true if there is scopes, false otherwise
func (obj *mainCommandInstruction) HasScopes() bool {
	return obj.scopes != nil
}

// Scopes retruns the scopes, if any
func (obj *mainCommandInstruction) Scopes() Scopes {
	return obj.scopes
}
