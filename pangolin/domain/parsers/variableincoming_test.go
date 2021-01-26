package parsers

import (
	"testing"
)

func Test_variableIncoming_mandatory_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableIncoming", grammarFile)

	file := "./tests/codes/variableincoming/mandatory.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	incoming := ins.(VariableIncoming)
	if !incoming.IsMandatory() {
		t.Errorf("the variableIncoming was expected to be mandatory")
		return
	}

	if incoming.IsOptional() {
		t.Errorf("the variableIncoming was NOT expected to be optional")
		return
	}
}

func Test_variableIncoming_optional_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableIncoming", grammarFile)

	file := "./tests/codes/variableincoming/optional.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	incoming := ins.(VariableIncoming)
	if incoming.IsMandatory() {
		t.Errorf("the variableIncoming was NOT expected to be mandatory")
		return
	}

	if !incoming.IsOptional() {
		t.Errorf("the variableIncoming was expected to be optional")
		return
	}

	val := incoming.OptionalDefaultValue()
	if !val.IsNumeric() {
		t.Errorf("the variableIncoming was expected to contain a numeric optional value")
		return
	}

	num := val.Numeric()
	if !num.IsInt() {
		t.Errorf("the variableIncoming was expected to contain an int optional value")
		return
	}

	in := num.Int()
	if *in != 56 {
		t.Errorf("the variableIncoming was expected to contain an optional value of %d, %d returned", 56, *in)
		return
	}
}
