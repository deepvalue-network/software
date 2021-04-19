package parsers

import (
	"testing"
)

func Test_languageTestDeclaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageTestDeclaration", grammarFile)

	file := "./tests/codes/languagetestdeclaration/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(LanguageTestDeclaration)
	if decl.Name() != "myTestFunc" {
		t.Errorf("the name was expected to be %s, %s returned", "myTestFunc", decl.Name())
		return
	}

	list := decl.Instructions()
	if len(list) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(list))
		return
	}
}
