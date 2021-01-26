package parsers

import (
	"testing"
)

func Test_patternMatch_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("patternMatch", grammarFile)

	file := "./tests/codes/patternmatch/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	match := ins.(PatternMatch)
	if match.Pattern() != "myPattern" {
		t.Errorf("the pattern was expected to be %s, %s returned", "myPattern", match.Pattern())
		return
	}

	if match.Labels() == nil {
		t.Errorf("the patternLabels was expected to be valid")
		return
	}

	if match.Variable() != "MyTokenVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "MyTokenVariable", match.Variable())
		return
	}
}
