package parsers

import (
	"testing"
)

func Test_value_boolValue_true_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/boolvalue_true.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsBool() {
		t.Errorf("the value was expected to be boolean")
		return
	}

	p := val.Bool()
	if *p != true {
		t.Errorf("the value was expected to be true")
		return
	}
}

func Test_value_boolValue_false_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/boolvalue_false.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsBool() {
		t.Errorf("the value was expected to be boolean")
		return
	}

	p := val.Bool()
	if *p != false {
		t.Errorf("the value was expected to be false")
		return
	}
}

func Test_value_nil_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/nil.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsNil() {
		t.Errorf("the value was expected to be nil")
		return
	}
}

func Test_value_numericValue_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/numeric_value.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsNumeric() {
		t.Errorf("the value was expected to be a numeric value")
		return
	}
}

func Test_value_stringValue_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/string_value.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsString() {
		t.Errorf("the value was expected to be a string value")
		return
	}
}

func Test_value_variable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("value", grammarFile)

	file := "./tests/codes/value/variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(Value)
	if !val.IsVariable() {
		t.Errorf("the value was expected to be a variable value")
		return
	}
}
