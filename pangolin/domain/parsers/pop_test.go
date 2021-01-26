package parsers

import (
	"testing"
)

func Test_pop_simple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("pop", grammarFile)

	file := "./tests/codes/pop/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	pop := ins.(Pop)
	if pop.HasStackFrame() {
		t.Errorf("the pop was not expecting a StackFrame")
		return
	}
}

func Test_pop_simple_withTransformOperation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("pop", grammarFile)

	file := "./tests/codes/pop/with_transformoperation.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	pop := ins.(Pop)
	if !pop.HasStackFrame() {
		t.Errorf("the pop was expecting a StackFrame")
		return
	}
}
