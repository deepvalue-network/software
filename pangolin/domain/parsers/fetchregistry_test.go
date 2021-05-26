package parsers

import (
	"testing"
)

func Test_fetchRegistry_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("fetchRegistry", grammarFile)

	file := "./tests/codes/fetchregistry/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(FetchRegistry)
	if reg.To() != "to" {
		t.Errorf("the to variable was expected to be %s, %s returned", "to", reg.To())
		return
	}

	if reg.From() != "from" {
		t.Errorf("the from variable was expected to be %s, %s returned", "to", reg.From())
		return
	}

	if reg.HasIndex() {
		t.Errorf("the register was NOT expected to contain an index")
		return
	}
}

func Test_fetchRegistry_withIndex_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("fetchRegistry", grammarFile)

	file := "./tests/codes/fetchregistry/with_index.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(FetchRegistry)
	if reg.To() != "to" {
		t.Errorf("the to variable was expected to be %s, %s returned", "to", reg.To())
		return
	}

	if reg.From() != "from" {
		t.Errorf("the from variable was expected to be %s, %s returned", "to", reg.From())
		return
	}

	if !reg.HasIndex() {
		t.Errorf("the register was expected to contain an index")
		return
	}
}
