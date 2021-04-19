package parsers

import (
	"testing"
)

func Test_languageLabelSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageLabelSection", grammarFile)

	file := "./tests/codes/languagelabelsection/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(LanguageLabelSection)
	list := section.Declarations()
	if len(list) != 2 {
		t.Errorf("%d declarations were expected, %d returned", 2, len(list))
		return
	}
}
