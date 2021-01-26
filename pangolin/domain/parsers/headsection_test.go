package parsers

import (
	"testing"
)

func Test_headSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headSection", grammarFile)

	file := "./tests/codes/headsection/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(HeadSection)
	if section.Name() != "my_name" {
		t.Errorf("the name was expected to be %s, %s returned", "my_name", section.Name())
		return
	}

	if section.Version() != "2020.03.04" {
		t.Errorf("the version was expected to be %s, %s returned", "2020.03.04", section.Version())
		return
	}

	if section.HasImport() {
		t.Errorf("the HeadSection was NOT expecting imports")
		return
	}
}

func Test_headSection_withImport_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headSection", grammarFile)

	file := "./tests/codes/headsection/with_import.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(HeadSection)
	if section.Name() != "my_name" {
		t.Errorf("the name was expected to be %s, %s returned", "my_name", section.Name())
		return
	}

	if section.Version() != "2020.03.04" {
		t.Errorf("the version was expected to be %s, %s returned", "2020.03.04", section.Version())
		return
	}

	if !section.HasImport() {
		t.Errorf("the HeadSection was expecting imports")
		return
	}
}
