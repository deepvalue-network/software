package parsers

import (
	"testing"
)

func Test_testCommand_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testCommand", grammarFile)

	file := "./tests/codes/testcommand/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(TestCommand)
	if cmd.Name() != "myTest" {
		t.Errorf("the test name was expected to be %s, %s returned", "myTest", cmd.Name())
		return
	}

	if cmd.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", cmd.Variable())
		return
	}

	list := cmd.Instructions()
	if len(list) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(list))
		return
	}
}
