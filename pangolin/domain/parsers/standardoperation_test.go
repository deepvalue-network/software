package parsers

import (
	"testing"
)

func Test_standardOperation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("standardOperation", grammarFile)

	file := "./tests/codes/standardoperation/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	operation := ins.(StandardOperation)
	if operation.First() != "myFirst" {
		t.Errorf("the first identifier was expected to be %s, %s returned", "myFirst", operation.First())
		return
	}

	if operation.Second() != "myOtherVariable" {
		t.Errorf("the first variable was expected to be %s, %s returned", "myOtherVariable", operation.First())
		return
	}

	if operation.Result() != "myVariable" {
		t.Errorf("the result variable was expected to be %s, %s returned", "myVariable", operation.Result())
		return
	}
}
