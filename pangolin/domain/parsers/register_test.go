package parsers

import (
	"testing"
)

func Test_register_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("register", grammarFile)

	file := "./tests/codes/register/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Register)
	if reg.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", reg.Variable())
		return
	}

	if reg.HasIndex() {
		t.Errorf("the register was NOT expected to contain an index")
		return
	}
}

func Test_register_withIndex_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("register", grammarFile)

	file := "./tests/codes/register/with_index.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Register)
	if reg.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", reg.Variable())
		return
	}

	if !reg.HasIndex() {
		t.Errorf("the register was expected to contain an index")
		return
	}
}
