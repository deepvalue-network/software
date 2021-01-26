package parsers

import (
	"testing"
)

func Test_variableDeclaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDeclaration", grammarFile)

	file := "./tests/codes/variabledeclaration/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(VariableDeclaration)
	if decl.Variable() != "MyVariable" {
		t.Errorf("the VariableDeclaration was expected %s as variable, %s returned", "MyVariable", decl.Variable())
		return
	}

	if decl.HasDirection() {
		t.Errorf("the declaration was not expecting a direction")
		return
	}
}

func Test_variableDeclaration_withDirection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDeclaration", grammarFile)

	file := "./tests/codes/variabledeclaration/with_direction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(VariableDeclaration)
	if decl.Variable() != "MyVariable" {
		t.Errorf("the VariableDeclaration was expected %s as variable, %s returned", "MyVariable", decl.Variable())
		return
	}

	if !decl.HasDirection() {
		t.Errorf("the declaration was expecting a direction")
		return
	}
}
