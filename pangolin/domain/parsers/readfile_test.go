package parsers

import (
	"testing"
)

func Test_readFile_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("readFile", grammarFile)

	file := "./tests/codes/readfile/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	readFile := ins.(ReadFile)
	if readFile == nil {
		t.Errorf("the ReadFile was expected to be valid")
	}
}
