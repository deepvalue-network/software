package parsers

import (
	"testing"
)

func Test_concatenation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("concatenation", grammarFile)

	file := "./tests/codes/concatenation/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	concat := ins.(Concatenation)
	op := concat.Operation()
	if op.Result().Local() != "answer" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "answer", op.Result().Local())
		return
	}

	if op.First().Variable().Local() != "first" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "first", op.Result().Local())
		return
	}

	if op.Second().Variable().Local() != "second" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "second", op.Result().Local())
		return
	}
}
