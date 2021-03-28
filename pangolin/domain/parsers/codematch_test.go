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
	if code.Content() != "myReturn" {
		t.Errorf("the content variable was expected to be %s, %s returned", "myReturn", code.Content())
		return
	}

	if code.Section() != "mySection" {
		t.Errorf("the content variable was expected to be %s, %s returned", "mySection", code.Section())
		return
	}

	patternVariables := code.PatternVariables()
	if len(patternVariables) != 2 {
		t.Errorf("%d pattern variables were expected, %d returned", 2, len(patternVariables))
		return
	}
}
