package parsers

import (
	"testing"
)

func Test_constantSection_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("constantSection", grammarFile)

	file := "./tests/codes/constantsection/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(ConstantSection)
	lst := section.Declarations()
	if len(lst) != 1 {
		t.Errorf("the ConstantSection was expecting %d declarations, %d returned", 1, len(lst))
		return
	}
}

func Test_constantSection_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("constantSection", grammarFile)

	file := "./tests/codes/constantsection/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(ConstantSection)
	lst := section.Declarations()
	if len(lst) != 3 {
		t.Errorf("the ConstantSection was expecting %d declarations, %d returned", 3, len(lst))
		return
	}
}
