package parsers

import (
	"testing"
)

func Test_assignment_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("assignment", grammarFile)

	file := "./tests/codes/assignment/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ass := ins.(Assignment)
	if ass.Variable() != "myVariable" {
		t.Errorf("the Assignment variable was expected to be %s, %s returned", "myVariable", ass.Variable())
		return
	}

	if !ass.Value().Numeric().IsInt() {
		t.Errorf("the Assignment value was expected to be an int")
		return
	}
}
