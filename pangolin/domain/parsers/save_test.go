package parsers

import (
	"testing"
)

func Test_save_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("save", grammarFile)

	file := "./tests/codes/save/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	save := ins.(Save)
	if save.To() != "to" {
		t.Errorf("the to variable was expected to be %s, %s returned", "to", save.To())
		return
	}

	if save.HasFrom() {
		t.Errorf("the from variable was NOT expected to be valid")
		return
	}
}

func Test_save_withFrom_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("save", grammarFile)

	file := "./tests/codes/save/with_from.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	save := ins.(Save)
	if save.To() != "to" {
		t.Errorf("the to variable was expected to be %s, %s returned", "to", save.To())
		return
	}

	if !save.HasFrom() {
		t.Errorf("the from variable was expected to be valid")
		return
	}

	if save.From() != "from" {
		t.Errorf("the from variable was expected to be %s, %s returned", "from", save.From())
		return
	}
}
