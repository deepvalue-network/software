package parsers

import (
	"testing"
)

func Test_language_withApplication_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/application.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if !lang.IsApplication() {
		t.Errorf("the language was expecting an application")
		return
	}

	if lang.IsDefinition() {
		t.Errorf("the language was NOT expecting a definition")
		return
	}
}

func Test_language_withDefinition_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/definition.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if lang.IsApplication() {
		t.Errorf("the language was NOT expecting an application")
		return
	}

	if !lang.IsDefinition() {
		t.Errorf("the language was expecting a definition")
		return
	}
}
