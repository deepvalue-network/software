package parsers

import (
	"testing"
)

func Test_mainCommandInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("mainCommandInstruction", grammarFile)

	file := "./tests/codes/maincommandinstruction/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(MainCommandInstruction)
	if cmd.HasScopes() {
		t.Errorf("the mainCommandInstruction was NOT expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}

func Test_mainCommandInstruction_withScopes_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("mainCommandInstruction", grammarFile)

	file := "./tests/codes/maincommandinstruction/with_scopes.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(MainCommandInstruction)
	if !cmd.HasScopes() {
		t.Errorf("the mainCommandInstruction was expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}
