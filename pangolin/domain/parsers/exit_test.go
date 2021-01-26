package parsers

import (
	"testing"
)

func Test_exit_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("exit", grammarFile)

	file := "./tests/codes/exit/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ex := ins.(Exit)
	if ex.HasCondition() {
		t.Errorf("the exit was not expecting a condition")
		return
	}
}

func Test_exit_withCondition_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("exit", grammarFile)

	file := "./tests/codes/exit/with_condition.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ex := ins.(Exit)
	if !ex.HasCondition() {
		t.Errorf("the exit was expecting a condition")
		return
	}
}
