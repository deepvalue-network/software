package parsers

import (
	"testing"
)

func Test_specificTokenCode_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("specificTokenCode", grammarFile)

	file := "./tests/codes/specifictokencode/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(SpecificTokenCode)
	if code.VariableName().String() != "myReturnVariable" {
		t.Errorf("the content variable was expected to be %s, %s returned", "myReturnVariable", code.VariableName().String())
		return
	}

	if code.PatternVariable() != "myToken" {
		t.Errorf("the tokenVariable was expected to be %s, %s returned", "myToken", code.PatternVariable())
		return
	}

	if code.HasAmount() {
		t.Errorf("the SpecificTokenCode was NOT expected to contain an amount")
		return
	}
}
