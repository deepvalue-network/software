package parsers

import (
	"testing"
)

func Test_labelCommand_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelCommand", grammarFile)

	file := "./tests/codes/labelcommand/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(LabelCommand)
	if cmd.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", cmd.Variable())
		return
	}

	if cmd.Name() != "myLabel" {
		t.Errorf("the variable was expected to be %s, %s returned", "myLabel", cmd.Name())
		return
	}

	list := cmd.Instructions()
	if len(list) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(list))
		return
	}
}
