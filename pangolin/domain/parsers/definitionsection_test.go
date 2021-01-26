package parsers

import (
	"testing"
)

func Test_definition_withConstants_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("definitionSection", grammarFile)

	file := "./tests/codes/definitionsection/with_constants.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	def := ins.(DefinitionSection)
	if !def.HasConstants() {
		t.Errorf("the definitionSection was expected to have constants")
		return
	}

	if def.HasVariables() {
		t.Errorf("the definitionSection was not expected to have variables")
		return
	}
}

func Test_definition_withVariables_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("definitionSection", grammarFile)

	file := "./tests/codes/definitionsection/with_variables.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	def := ins.(DefinitionSection)
	if def.HasConstants() {
		t.Errorf("the definitionSection was not expected to have constants")
		return
	}

	if !def.HasVariables() {
		t.Errorf("the definitionSection was expected to have variables")
		return
	}
}

func Test_definition_withConstants_withVariables_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("definitionSection", grammarFile)

	file := "./tests/codes/definitionsection/with_constants_with_variables.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	def := ins.(DefinitionSection)
	if !def.HasConstants() {
		t.Errorf("the definitionSection was expected to have constants")
		return
	}

	if !def.HasVariables() {
		t.Errorf("the definitionSection was expected to have variables")
		return
	}
}
