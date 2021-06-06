package parsers

import (
	"testing"
)

func Test_module_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("module", grammarFile)

	file := "./tests/codes/module/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	module := ins.(Module)
	if module.StackFrame() != "stackFrame" {
		t.Errorf("the stackFrame was expected to contain %s, %s returned", "stackFrame", module.StackFrame())
		return
	}

	if module.Name() != "module_name" {
		t.Errorf("the name was expected to contain %s, %s returned", "module_name", module.Name())
		return
	}

	if module.Symbol() != "symbol_name" {
		t.Errorf("the symbol was expected to contain %s, %s returned", "symbol_name", module.Symbol())
		return
	}
}
