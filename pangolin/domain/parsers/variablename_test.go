package parsers

import (
	"testing"
)

func Test_variableName_global_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableName", grammarFile)

	file := "./tests/codes/variablename/global.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	varName := ins.(VariableName)
	if !varName.IsGlobal() {
		t.Errorf("the variableName was expected to be global")
		return
	}

	if varName.Global() != "ThisIsGlobal" {
		t.Errorf("the variableName was expected to be %s, %s returned", "ThisIsGlobal", varName.Global())
		return
	}
}

func Test_variableName_local_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableName", grammarFile)

	file := "./tests/codes/variablename/local.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	varName := ins.(VariableName)
	if !varName.IsLocal() {
		t.Errorf("the variableName was expected to be local")
		return
	}

	if varName.Local() != "thisIsLocal" {
		t.Errorf("the variableName was expected to be %s, %s returned", "thisIsLocal", varName.Local())
		return
	}
}
