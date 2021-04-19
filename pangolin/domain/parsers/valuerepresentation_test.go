package parsers

import (
	"testing"
)

func Test_variableRepresentation_withVariable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("valueRepresentation", grammarFile)

	file := "./tests/codes/valuerepresentation/variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(ValueRepresentation)
	if !val.IsVariable() {
		t.Errorf("the ValueRepresentation was expected to be a variable")
		return
	}

	if val.IsValue() {
		t.Errorf("the ValueRepresentation was NOT expected to be a value")
		return
	}
}

func Test_variableRepresentation_withValue_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("valueRepresentation", grammarFile)

	file := "./tests/codes/valuerepresentation/value.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(ValueRepresentation)
	if val.IsVariable() {
		t.Errorf("the ValueRepresentation was NOT expected to be a variable")
		return
	}

	if !val.IsValue() {
		t.Errorf("the ValueRepresentation was expected to be a value")
		return
	}
}
