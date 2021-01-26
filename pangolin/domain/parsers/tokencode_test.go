package parsers

import (
	"testing"
)

func Test_tokenCode_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("tokenCode", grammarFile)

	file := "./tests/codes/tokencode/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(TokenCode)
	if code.Content().String() != "myReturnVariable" {
		t.Errorf("the content variable was expected to be %s, %s returned", "myReturnVariable", code.Content().String())
		return
	}

	if code.TokenVariable() != "MyTokenVariable" {
		t.Errorf("the tokenVariable was expected to be %s, %s returned", "MyTokenVariable", code.TokenVariable())
		return
	}
}
