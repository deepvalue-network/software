package parsers

import (
	"testing"
)

func Test_scriptTests_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptTests", grammarFile)

	file := "./tests/codes/scripttests/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	tests := ins.(ScriptTests)
	list := tests.All()
	if len(list) != 1 {
		t.Errorf("%d elements were expected in the ScriptTests list, %d returned", 1, len(list))
		return
	}
}

func Test_scriptTests_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptTests", grammarFile)

	file := "./tests/codes/scripttests/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	tests := ins.(ScriptTests)
	list := tests.All()
	if len(list) != 2 {
		t.Errorf("%d elements were expected in the ScriptTests list, %d returned", 2, len(list))
		return
	}
}
