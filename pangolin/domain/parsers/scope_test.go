package parsers

import (
	"testing"
)

func Test_scope_internal_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scope", grammarFile)

	file := "./tests/codes/scope/internal.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	scope := ins.(Scope)
	if !scope.IsInternal() {
		t.Errorf("the scope was expected to be internal")
		return
	}

	if scope.IsExternal() {
		t.Errorf("the scope was NOT expected to be external")
		return
	}
}

func Test_scope_external_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scope", grammarFile)

	file := "./tests/codes/scope/external.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	scope := ins.(Scope)
	if scope.IsInternal() {
		t.Errorf("the scope was NOT expected to be internal")
		return
	}

	if !scope.IsExternal() {
		t.Errorf("the scope was expected to be external")
		return
	}
}
