package parsers

import (
	"testing"
)

func Test_scope_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scopes", grammarFile)

	file := "./tests/codes/scopes/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	scopes := ins.(Scopes)
	list := scopes.All()
	if len(list) != 1 {
		t.Errorf("%d Scope instances were expected, %d returned", 1, len(list))
		return
	}
}

func Test_scope_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scopes", grammarFile)

	file := "./tests/codes/scopes/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	scopes := ins.(Scopes)
	list := scopes.All()
	if len(list) != 5 {
		t.Errorf("%d Scope instances were expected, %d returned", 5, len(list))
		return
	}
}
