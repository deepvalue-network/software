package parsers

import (
	"testing"
)

func Test_switch_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("switch", grammarFile)

	file := "./tests/codes/switch/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	swtch := ins.(Switch)
	if swtch.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", swtch.Variable())
		return
	}
}
