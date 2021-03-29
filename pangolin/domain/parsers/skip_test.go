package parsers

import (
	"testing"
)

func Test_skip_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("skip", grammarFile)

	file := "./tests/codes/skip/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	skip := ins.(Skip)
	if skip.Pointer() == nil {
		t.Errorf("the IntPointer inside the Skip instance was expected to be valid")
		return
	}
}
