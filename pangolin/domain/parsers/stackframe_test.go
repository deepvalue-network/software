package parsers

import (
	"testing"
)

func Test_stackframe_pop_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("stackFrame", grammarFile)

	file := "./tests/codes/stackframe/pop.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	frame := ins.(StackFrame)
	if !frame.IsPop() {
		t.Errorf("the stackFrame was expected to be a pop instruction")
		return
	}
}

func Test_stackframe_push_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("stackFrame", grammarFile)

	file := "./tests/codes/stackframe/push.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	frame := ins.(StackFrame)
	if !frame.IsPush() {
		t.Errorf("the stackFrame was expected to be a push instruction")
		return
	}
}

func Test_stackframe_WithSkip_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("stackFrame", grammarFile)

	file := "./tests/codes/stackframe/skip.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	frame := ins.(StackFrame)
	if !frame.IsSkip() {
		t.Errorf("the stackFrame was expected to be a skip instruction")
		return
	}
}

func Test_stackframe_WithIndex_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("stackFrame", grammarFile)

	file := "./tests/codes/stackframe/index.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	frame := ins.(StackFrame)
	if !frame.IsIndex() {
		t.Errorf("the stackFrame was expected to be an index instruction")
		return
	}
}
