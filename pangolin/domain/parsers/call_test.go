package parsers

import (
	"testing"
)

func Test_call_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("call", grammarFile)

	file := "./tests/codes/call/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	call := ins.(Call)
	if call.Name() != "my_imported_app" {
		t.Errorf("the call name was expected to be: %s, %s returned", "my_imported_app", call.Name())
		return
	}

	if call.StackFrame() != "myStackFrame" {
		t.Errorf("the call stackFrame was expected to be: %s, %s returned", "myStacKFrame", call.StackFrame())
		return
	}

	if call.HasCondition() {
		t.Errorf("the call was not expecting a condition")
		return
	}
}

func Test_call_withCondition_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("call", grammarFile)

	file := "./tests/codes/call/with_condition.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	call := ins.(Call)
	if call.Name() != "my_imported_app" {
		t.Errorf("the call name was expected to be: %s, %s returned", "my_imported_app", call.Name())
		return
	}

	if call.StackFrame() != "myStackFrame" {
		t.Errorf("the call stackFrame was expected to be: %s, %s returned", "myStacKFrame", call.StackFrame())
		return
	}

	if !call.HasCondition() {
		t.Errorf("the call was expecting a condition")
		return
	}
}
