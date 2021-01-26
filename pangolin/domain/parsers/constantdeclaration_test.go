package parsers

import (
	"testing"
)

func Test_constantDeclaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("constantDeclaration", grammarFile)

	file := "./tests/codes/constantdeclaration/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(ConstantDeclaration)
	if decl.Constant() != "MY_CONSTANT" {
		t.Errorf("the VariableDeclaration was expected %s as variable, %s returned", "MY_CONSTANT", decl.Constant())
		return
	}
}
