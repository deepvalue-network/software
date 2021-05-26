package parsers

import (
	"testing"
)

func Test_unregister_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("unregister", grammarFile)

	file := "./tests/codes/unregister/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Unregister)
	if reg.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", reg.Variable())
		return
	}
}
