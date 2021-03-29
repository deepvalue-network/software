package parsers

import (
	"testing"
)

func Test_index_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("index", grammarFile)

	file := "./tests/codes/index/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	index := ins.(Index)
	if index.Variable() != "myVariable" {
		t.Errorf("the variable in the Index was expected to be %s, %s returned", "myVariable", index.Variable())
		return
	}
}
