package parsers

import (
	"testing"
)

func Test_languageApplication_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageApplication", grammarFile)

	file := "./tests/codes/languageapplication/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(LanguageApplication)
	if app.HasTests() {
		t.Errorf("no tests were expected in the LanguageApplication")
		return
	}
}

func Test_languageApplication_withTests_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageApplication", grammarFile)

	file := "./tests/codes/languageapplication/with_tests.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(LanguageApplication)
	if !app.HasTests() {
		t.Errorf("tests were expected in the LanguageApplication")
		return
	}
}
