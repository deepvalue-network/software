package parsers

type testCommandInstruction struct {
	ins    TestInstruction
	scopes Scopes
}

func createTestCommandInstruction(
	ins TestInstruction,
) TestCommandInstruction {
	return createTestCommandInstructionInternally(ins, nil)
}

func createTestCommandInstructionWithScopes(
	ins TestInstruction,
	scopes Scopes,
) TestCommandInstruction {
	return createTestCommandInstructionInternally(ins, scopes)
}

func createTestCommandInstructionInternally(
	ins TestInstruction,
	scopes Scopes,
) TestCommandInstruction {
	out := testCommandInstruction{
		ins:    ins,
		scopes: scopes,
	}

	return &out
}

// Instruction returns the instruction
func (obj *testCommandInstruction) Instruction() TestInstruction {
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
