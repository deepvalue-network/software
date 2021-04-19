package parsers

import (
	"testing"
)

func Test_languageLabelDeclaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageLabelDeclaration", grammarFile)

	file := "./tests/codes/languagelabeldeclaration/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(LanguageLabelDeclaration)
	if decl.Name() != "myLabel" {
		t.Errorf("the label name was expected to be %s, %s returned", "myLabel", decl.Name())
		return
	}

	list := decl.Instructions()
	if len(list) != 2 {
		t.Errorf("%d instructions were expected, %d returned", 2, len(list))
		return
	}
}
