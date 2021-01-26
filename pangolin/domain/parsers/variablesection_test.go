package parsers

import (
	"testing"
)

func Test_variableSection_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableSection", grammarFile)

	file := "./tests/codes/variablesection/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(VariableSection)
	lst := section.Declarations()
	if len(lst) != 1 {
		t.Errorf("the VariableSection was expecting %d declarations, %d returned", 1, len(lst))
		return
	}
}

func Test_variableSection_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableSection", grammarFile)

	file := "./tests/codes/variablesection/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(VariableSection)
	lst := section.Declarations()
	if len(lst) != 3 {
		t.Errorf("the VariableSection was expecting %d declarations, %d returned", 3, len(lst))
		return
	}
}
