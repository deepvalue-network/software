package parsers

import (
	"testing"
)

func Test_logical_and_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("logical", grammarFile)

	file := "./tests/codes/logical/and.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	logical := ins.(Logical)
	if !logical.IsAnd() {
		t.Errorf("the logical operator was expected to be an AND")
		return
	}
}

func Test_logical_or_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("logical", grammarFile)

	file := "./tests/codes/logical/or.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	logical := ins.(Logical)
	if !logical.IsOr() {
		t.Errorf("the logical operator was expected to be an OR")
		return
	}
}
