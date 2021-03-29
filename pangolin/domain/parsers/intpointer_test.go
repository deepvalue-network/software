package parsers

import (
	"testing"
)

func Test_intPointer_withInt_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("intPointer", grammarFile)

	file := "./tests/codes/intpointer/int.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	pointer := ins.(IntPointer)
	if !pointer.IsInt() {
		t.Errorf("the IntPointer was expected to contain an Int")
		return
	}

	if pointer.IsVariable() {
		t.Errorf("the IntPointer was NOT expected to contain a variable")
		return
	}

	if pointer.Int() != 56 {
		t.Errorf("the IntPointer was expected to contain %d, %d returned", 56, pointer.Int())
		return
	}
}

func Test_intPointer_withVariable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("intPointer", grammarFile)

	file := "./tests/codes/intpointer/variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	pointer := ins.(IntPointer)
	if pointer.IsInt() {
		t.Errorf("the IntPointer was NOT expected to contain an Int")
		return
	}

	if !pointer.IsVariable() {
		t.Errorf("the IntPointer was expected to contain a variable")
		return
	}

	if pointer.Variable() != "myVariable" {
		t.Errorf("the IntPointer was expected to contain %s, %s returned", "myVariable", pointer.Variable())
		return
	}
}
