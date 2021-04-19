package parsers

import (
	"testing"
)

func Test_testCommandInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testCommandInstruction", grammarFile)

	file := "./tests/codes/testcommandinstruction/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(TestCommandInstruction)
	if cmd.HasScopes() {
		t.Errorf("the testCommandInstruction was NOT expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}

func Test_testCommandInstruction_withScopes_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testCommandInstruction", grammarFile)

	file := "./tests/codes/testcommandinstruction/with_scopes.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(TestCommandInstruction)
	if !cmd.HasScopes() {
		t.Errorf("the testCommandInstruction was expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}
