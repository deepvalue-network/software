package parsers

import (
	"testing"
)

func Test_transformOperation_withConstant_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("transformOperation", grammarFile)

	file := "./tests/codes/transformoperation/with_constant.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	operation := ins.(TransformOperation)
	if operation.Input().Constant() != "MY_CONSTANT" {
		t.Errorf("the identifier was expected to be %s, %s returned", "MY_CONSTANT", operation.Input().Constant())
		return
	}

	if operation.Result().Local() != "myVariable" {
		t.Errorf("the result variable was expected to be %s, %s returned", "myVariable", operation.Result().Local())
		return
	}
}

func Test_transformOperation_withVariable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("transformOperation", grammarFile)

	file := "./tests/codes/transformoperation/with_variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	operation := ins.(TransformOperation)
	if operation.Input().Variable().Local() != "myOtherVariable" {
		t.Errorf("the identifier was expected to be %s, %s returned", "myOtherVariable", operation.Input().Variable().Local())
		return
	}

	if operation.Result().Local() != "myVariable" {
		t.Errorf("the result variable was expected to be %s, %s returned", "myVariable", operation.Result().Local())
		return
	}
}
