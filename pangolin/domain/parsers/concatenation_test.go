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
	if op.Result() != "answer" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "answer", op.Result())
		return
	}

	if op.First() != "first" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "first", op.Result())
		return
	}

	if op.Second() != "second" {
		t.Errorf("the Concatenation was expected '%s' as result, %s returned", "second", op.Result())
		return
	}
}
