package parsers

import (
	"testing"
)

func Test_languageMainSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageMainSection", grammarFile)

	file := "./tests/codes/languagemainsection/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(LanguageMainSection)
	list := section.Instructions()
	if len(list) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(list))
		return
	}
}
