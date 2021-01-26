package parsers

import (
	"testing"
)

func Test_match_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("match", grammarFile)

	file := "./tests/codes/match/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	match := ins.(Match)
	input := match.Input()
	if input.String() != "myInput" {
		t.Errorf("the input Identifier was expected to be %s, %s returned", "myInput", input.String())
		return
	}

	if match.HasPattern() {
		t.Errorf("the pattern was expected to NOT be valid")
		return
	}
}

func Test_match_withPattern_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("match", grammarFile)

	file := "./tests/codes/match/with_pattern.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	match := ins.(Match)
	input := match.Input()
	if input.String() != "myInput" {
		t.Errorf("the input Identifier was expected to be %s, %s returned", "myInput", input.String())
		return
	}

	if !match.HasPattern() {
		t.Errorf("the pattern was expected to be valid")
		return
	}

	pattern := match.Pattern()
	if pattern != "myPattern" {
		t.Errorf("the pattern was expected to be %s, %s returned", "myPattern", pattern)
		return
	}
}
