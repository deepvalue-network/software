package parsers

import (
	"testing"
)

func Test_codeMatch_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("codeMatch", grammarFile)

	file := "./tests/codes/codematch/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(CodeMatch)
	if code.Content().String() != "myReturn" {
		t.Errorf("the content variable was expected to be %s, %s returned", "myReturn", code.Content().String())
		return
	}

	if code.Section().String() != "mySection" {
		t.Errorf("the content variable was expected to be %s, %s returned", "mySection", code.Section().String())
		return
	}

	if code.TokenVariable() != "MyTokenVariable" {
		t.Errorf("the token variable was expected to be %s, %s returned", "MyTokenVariable", code.TokenVariable())
		return
	}

	patternVariables := code.PatternVariables()
	if len(patternVariables) != 2 {
		t.Errorf("%d pattern variables were expected, %d returned", 2, len(patternVariables))
		return
	}
}
