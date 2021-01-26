package parsers

import (
	"testing"
)

func Test_testSection_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testSection", grammarFile)

	file := "./tests/codes/testsection/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(TestSection)
	lst := section.Declarations()
	if len(lst) != 1 {
		t.Errorf("the testSection was expecting %d declarations, %d returned", 1, len(lst))
		return
	}
}

func Test_testSection_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testSection", grammarFile)

	file := "./tests/codes/testsection/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(TestSection)
	lst := section.Declarations()
	if len(lst) != 2 {
		t.Errorf("the testSection was expecting %d declarations, %d returned", 2, len(lst))
		return
	}
}
