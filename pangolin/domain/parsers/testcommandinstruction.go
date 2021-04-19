package parsers

type testCommandInstruction struct {
	ins    LanguageTestInstruction
	scopes Scopes
}

func createTestCommandInstruction(
	ins LanguageTestInstruction,
) TestCommandInstruction {
	return createTestCommandInstructionInternally(ins, nil)
}

func createTestCommandInstructionWithScopes(
	ins LanguageTestInstruction,
	scopes Scopes,
) TestCommandInstruction {
	return createTestCommandInstructionInternally(ins, scopes)
}

func createTestCommandInstructionInternally(
	ins LanguageTestInstruction,
	scopes Scopes,
) TestCommandInstruction {
	out := testCommandInstruction{
		ins:    ins,
		scopes: scopes,
	}

	return &out
}

// Instruction returns the instruction
func (obj *testCommandInstruction) Instruction() LanguageTestInstruction {
	return obj.ins
}

// HasScopes returns true if there is scopes, false otherwise
func (obj *testCommandInstruction) HasScopes() bool {
	return obj.scopes != nil
}

// Scopes returns the scopes, if any
func (obj *testCommandInstruction) Scopes() Scopes {
	return obj.scopes
}
