package parsers

import (
	"testing"
)

func Test_relational_lessThan_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relational", grammarFile)

	file := "./tests/codes/relational/less_than.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	rel := ins.(Relational)
	if !rel.IsLessThan() {
		t.Errorf("the relational operator was expected to be a LessThan")
		return
	}
}

func Test_relational_equal_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relational", grammarFile)

	file := "./tests/codes/relational/equal.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	rel := ins.(Relational)
	if !rel.IsEqual() {
		t.Errorf("the relational operator was expected to be an Equal")
		return
	}
}

func Test_relational_notEqual_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relational", grammarFile)

	file := "./tests/codes/relational/not_equal.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	rel := ins.(Relational)
	if !rel.IsNotEqual() {
		t.Errorf("the relational operator was expected to be a NotEqual")
		return
	}
}
