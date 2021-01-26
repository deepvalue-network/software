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

func Test_stackframe_frameAssignment_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("stackFrame", grammarFile)

	file := "./tests/codes/stackframe/frameassignment.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	frame := ins.(StackFrame)
	if !frame.IsAssignment() {
		t.Errorf("the stackFrame was expected to be a frameAssignment instruction")
		return
	}
}
