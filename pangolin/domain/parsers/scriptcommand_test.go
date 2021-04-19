package parsers

import (
	"testing"
)

func Test_scriptCommand_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptCommand", grammarFile)

	file := "./tests/codes/scriptcommand/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(ScriptCommand)
	if cmd.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", cmd.Variable())
		return
	}

	list := cmd.Values()
	if len(list) != 2 {
		t.Errorf("%d values ([]ScriptValue) were expected, %d returned", 2, len(list))
		return
	}
}
