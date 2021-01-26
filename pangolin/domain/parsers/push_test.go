package parsers

import (
	"testing"
)

func Test_push_simple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("push", grammarFile)

	file := "./tests/codes/push/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	push := ins.(Push)
	if push.HasStackFrame() {
		t.Errorf("the push was not expecting a StackFrame")
		return
	}
}

func Test_push_simple_withVariableName_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("push", grammarFile)

	file := "./tests/codes/push/with_variablename.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	push := ins.(Push)
	if !push.HasStackFrame() {
		t.Errorf("the push was expecting a StackFrame")
		return
	}
}
