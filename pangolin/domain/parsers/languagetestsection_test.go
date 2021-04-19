package parsers

import (
	"testing"
)

func Test_languageTestSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageTestSection", grammarFile)

	file := "./tests/codes/languagetestsection/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(LanguageTestSection)
	list := section.Declarations()
	if len(list) != 2 {
		t.Errorf("%d declarations were expected, %d returned", 2, len(list))
		return
	}
}
