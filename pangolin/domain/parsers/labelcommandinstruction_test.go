package parsers

import (
	"testing"
)

func Test_labelCommandInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelCommandInstruction", grammarFile)

	file := "./tests/codes/labelcommandinstruction/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(LabelCommandInstruction)
	if cmd.HasScopes() {
		t.Errorf("the labelCommandInstruction was NOT expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}

func Test_labelCommandInstruction_withScopes_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelCommandInstruction", grammarFile)

	file := "./tests/codes/labelcommandinstruction/with_scopes.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(LabelCommandInstruction)
	if !cmd.HasScopes() {
		t.Errorf("the labelCommandInstruction was expected to contain scopes")
		return
	}

	if cmd.Instruction() == nil {
		t.Errorf("the instruction was expected to be valid")
		return
	}
}
