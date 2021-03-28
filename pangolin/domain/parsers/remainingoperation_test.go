package parsers

import (
	"testing"
)

func Test_remainingOperation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("remainingOperation", grammarFile)

	file := "./tests/codes/remainingoperation/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	operation := ins.(RemainingOperation)
	if operation.First().Constant() != "MY_CONSTANT" {
		t.Errorf("the first identifier was expected to be %s, %s returned", "MY_CONSTANT", operation.First().Constant())
		return
	}

	if operation.Second().Variable().Local() != "myOtherVariable" {
		t.Errorf("the first variable was expected to be %s, %s returned", "myOtherVariable", operation.First().Variable().Local())
		return
	}

	if operation.Result().Local() != "myVariable" {
		t.Errorf("the result variable was expected to be %s, %s returned", "myVariable", operation.Result().Local())
		return
	}

	if operation.Remaining().Local() != "myGlobalVariable" {
		t.Errorf("the remaining variable was expected to be %s, %s returned", "myGlobalVariable", operation.Remaining().Local())
		return
	}
}
