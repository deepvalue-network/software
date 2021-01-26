package parsers

import (
	"testing"
)

func Test_mainSection_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("mainSection", grammarFile)

	file := "./tests/codes/mainsection/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(MainSection)
	lst := section.Instructions()
	if len(lst) != 1 {
		t.Errorf("the MainSection was expected to contain %d instructions, %d returned", 1, len(lst))
		return
	}
}

func Test_mainSection_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("mainSection", grammarFile)

	file := "./tests/codes/mainsection/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(MainSection)
	lst := section.Instructions()
	if len(lst) != 4 {
		t.Errorf("the MainSection was expected to contain %d instructions, %d returned", 4, len(lst))
		return
	}
}
