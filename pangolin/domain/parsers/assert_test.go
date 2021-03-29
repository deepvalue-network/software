package parsers

import (
	"testing"
)

func Test_assert_simple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("assert", grammarFile)

	file := "./tests/codes/assert/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ass := ins.(Assert)
	if ass.Index() != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, ass.Index())
		return
	}

	if ass.HasCondition() {
		t.Errorf("the assert was expected to NOT contain a condition")
		return
	}
}

func Test_assert_withCondition_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("assert", grammarFile)

	file := "./tests/codes/assert/condition.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ass := ins.(Assert)
	if ass.Index() != 1 {
		t.Errorf("the index was expected to be %d, %d returned", 1, ass.Index())
		return
	}

	if !ass.HasCondition() {
		t.Errorf("the assert was expected to contain a condition")
		return
	}
}
