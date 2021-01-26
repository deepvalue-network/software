package parsers

import (
	"testing"
)

func Test_numericValue_int_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("numericValue", grammarFile)

	file := "./tests/codes/numericvalue/int.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(NumericValue)
	if val.IsNegative() {
		t.Errorf("the numericValue was expected to be positive")
		return
	}

	if !val.IsInt() {
		t.Errorf("the numericValue was expected to be an int")
		return
	}

	p := val.Int()
	if *p != 34 {
		t.Errorf("the numericValue was expected to be %d, %d returned", 34, *p)
		return
	}
}

func Test_numericValue_negative_int_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("numericValue", grammarFile)

	file := "./tests/codes/numericvalue/neg_int.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(NumericValue)
	if !val.IsNegative() {
		t.Errorf("the numericValue was expected to be negative")
		return
	}

	if !val.IsInt() {
		t.Errorf("the numericValue was expected to be an int")
		return
	}

	p := val.Int()
	if *p != 13 {
		t.Errorf("the numericValue was expected to be %d, %d returned", 13, *p)
		return
	}
}

func Test_numericValue_float_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("numericValue", grammarFile)

	file := "./tests/codes/numericvalue/float.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(NumericValue)
	if val.IsNegative() {
		t.Errorf("the numericValue was expected to be positive")
		return
	}

	if !val.IsFloat() {
		t.Errorf("the numericValue was expected to be a float")
		return
	}

	p := val.Float()
	if *p != 34.21 {
		t.Errorf("the numericValue was expected to be %f, %f returned", 34.21, *p)
		return
	}
}

func Test_numericValue_neg_float_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("numericValue", grammarFile)

	file := "./tests/codes/numericvalue/neg_float.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(NumericValue)
	if !val.IsNegative() {
		t.Errorf("the numericValue was expected to be negative")
		return
	}

	if !val.IsFloat() {
		t.Errorf("the numericValue was expected to be a float")
		return
	}

	p := val.Float()
	if *p != 56.356 {
		t.Errorf("the numericValue was expected to be %f, %f returned", 56.356, *p)
		return
	}
}
