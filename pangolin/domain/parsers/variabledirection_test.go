package parsers

import (
	"testing"
)

func Test_variableDirection_incoming_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDirection", grammarFile)

	file := "./tests/codes/variabledirection/in.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	dir := ins.(VariableDirection)
	if !dir.IsIncoming() {
		t.Errorf("the direction was expected to be incoming")
		return
	}

	if dir.IsOutgoing() {
		t.Errorf("the direction was not expected to be outgoing")
		return
	}
}

func Test_variableDirection_outgoing_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDirection", grammarFile)

	file := "./tests/codes/variabledirection/out.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	dir := ins.(VariableDirection)
	if dir.IsIncoming() {
		t.Errorf("the direction was not expected to be incoming")
		return
	}

	if !dir.IsOutgoing() {
		t.Errorf("the direction was expected to be outgoing")
		return
	}
}

func Test_variableDirection_incoming_outgoing_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDirection", grammarFile)

	file := "./tests/codes/variabledirection/in_out.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	dir := ins.(VariableDirection)
	if !dir.IsIncoming() {
		t.Errorf("the direction was expected to be incoming")
		return
	}

	if !dir.IsOutgoing() {
		t.Errorf("the direction was expected to be outgoing")
		return
	}
}

func Test_variableDirection_outgoin_incoming_g_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("variableDirection", grammarFile)

	file := "./tests/codes/variabledirection/out_in.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	dir := ins.(VariableDirection)
	if !dir.IsIncoming() {
		t.Errorf("the direction was expected to be incoming")
		return
	}

	if !dir.IsOutgoing() {
		t.Errorf("the direction was expected to be outgoing")
		return
	}
}
