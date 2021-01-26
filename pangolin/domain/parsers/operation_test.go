package parsers

import (
	"testing"
)

func Test_operation_arythmetic_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("operation", grammarFile)

	file := "./tests/codes/operation/arythmetic.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	op := ins.(Operation)
	if !op.IsArythmetic() {
		t.Errorf("the Operation was not expected to contain an Arythmetic")
		return
	}
}

func Test_operation_logical_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("operation", grammarFile)

	file := "./tests/codes/operation/logical.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	op := ins.(Operation)
	if !op.IsLogical() {
		t.Errorf("the Operation was not expected to contain a Logical")
		return
	}
}

func Test_operation_relational_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("operation", grammarFile)

	file := "./tests/codes/operation/relational.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	op := ins.(Operation)
	if !op.IsRelational() {
		t.Errorf("the Operation was not expected to contain a Relational")
		return
	}
}
