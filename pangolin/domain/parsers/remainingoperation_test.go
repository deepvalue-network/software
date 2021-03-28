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

	if operation.Remaining() != "myRemaining" {
		t.Errorf("the remaining variable was expected to be %s, %s returned", "myRemaining", operation.Remaining())
		return
	}
}
