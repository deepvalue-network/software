package parsers

import (
	"testing"
)

func Test_declaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("declaration", grammarFile)

	file := "./tests/codes/declaration/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(Declaration)
	if decl.Variable() != "myVariable" {
		t.Errorf("the Declaration was expected %s as variabkle, %s returned", "myVariable", decl.Variable())
		return
	}
}
