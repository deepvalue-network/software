package parsers

import (
	"testing"
)

func Test_variable_declaration_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variable", grammarFile)

	file := "./tests/codes/variable/declaration.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	variable := ins.(Variable)
	if !variable.IsDeclaration() {
		t.Errorf("the Variable was expected to be a declaration")
		return
	}
}

func Test_variable_assignment_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variable", grammarFile)

	file := "./tests/codes/variable/assignment.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	variable := ins.(Variable)
	if !variable.IsAssignment() {
		t.Errorf("the Variable was expected to be an assignment")
		return
	}
}

func Test_variable_concatenation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variable", grammarFile)

	file := "./tests/codes/variable/concatenation.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	variable := ins.(Variable)
	if !variable.IsConcatenation() {
		t.Errorf("the Variable was expected to be a concatenation")
		return
	}
}

func Test_variable_deleteVariable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variable", grammarFile)

	file := "./tests/codes/variable/delete_variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	variable := ins.(Variable)
	if !variable.IsDelete() {
		t.Errorf("the Variable was expected to be a delete variable")
		return
	}
}
