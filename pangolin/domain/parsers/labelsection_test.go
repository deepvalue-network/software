package parsers

import (
	"testing"
)

func Test_labelSection_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelSection", grammarFile)

	file := "./tests/codes/labelsection/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(LabelSection)
	lst := section.Declarations()
	if len(lst) != 1 {
		t.Errorf("the LabelSection was expected to contain %d declarations, %d returned", 1, len(lst))
		return
	}
}

func Test_labelSection_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelSection", grammarFile)

	file := "./tests/codes/labelsection/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(LabelSection)
	lst := section.Declarations()
	if len(lst) != 2 {
		t.Errorf("the LabelSection was expected to contain %d declarations, %d returned", 2, len(lst))
		return
	}
}
