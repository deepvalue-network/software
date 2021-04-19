package parsers

import (
	"testing"
)

func Test_languageCommand_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageCommand", grammarFile)

	file := "./tests/codes/languagecommand/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(LanguageCommand)
	if cmd.Variable() != "myVariable" {
		t.Errorf("the variable was expected to be %s, %s returned", "myVariable", cmd.Variable())
		return
	}

	list := cmd.Values()
	if len(list) != 2 {
		t.Errorf("%d values ([]LanguageValue) were expected, %d returned", 2, len(list))
		return
	}
}
