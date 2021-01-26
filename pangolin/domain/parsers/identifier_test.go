package parsers

import (
	"testing"
)

func Test_identifier_variableName_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("identifier", grammarFile)

	file := "./tests/codes/identifier/variablename.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	identifier := ins.(Identifier)
	if !identifier.IsVariable() {
		t.Errorf("the identifier was expected to be a variable")
		return
	}

	if identifier.Variable().Local() != "myName" {
		t.Errorf("the identifier was expected to be %s, %s returned", "myName", identifier.Variable().Local())
		return
	}
}

func Test_identifier_constant_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("identifier", grammarFile)

	file := "./tests/codes/identifier/constant.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	identifier := ins.(Identifier)
	if !identifier.IsConstant() {
		t.Errorf("the identifier was expected to be a constant")
		return
	}

	if identifier.Constant() != "THIS_IS_A_CONSTANT" {
		t.Errorf("the identifier was expected to be %s, %s returned", "THIS_IS_A_CONSTANT", identifier.Constant())
		return
	}
}
