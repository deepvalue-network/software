package parsers

import (
	"testing"
)

func Test_tokenSection_WithTokenCode_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("tokenSection", grammarFile)

	file := "./tests/codes/tokensection/tokencode.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(TokenSection)
	if !code.IsCode() {
		t.Errorf("the TokenSection was expected to contain a TokenCode instance")
		return
	}

	if code.IsSpecific() {
		t.Errorf("the TokenSection was NOT expected to contain a SpecificTokenCode instance")
		return
	}
}

func Test_tokenSection_WithSpecificTokenCode_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("tokenSection", grammarFile)

	file := "./tests/codes/tokensection/specifictokencode.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(TokenSection)
	if code.IsCode() {
		t.Errorf("the TokenSection was NOT expected to contain a TokenCode instance")
		return
	}

	if !code.IsSpecific() {
		t.Errorf("the TokenSection was expected to contain a SpecificTokenCode instance")
		return
	}

	if code.Specific().HasAmount() {
		t.Errorf("the TokenSection was expected to contain a SpecificTokenCode instance without an amount")
		return
	}
}

func Test_tokenSection_WithSpecificTokenCode_withAmount_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("tokenSection", grammarFile)

	file := "./tests/codes/tokensection/specifictokencodewithamount.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	code := ins.(TokenSection)
	if code.IsCode() {
		t.Errorf("the TokenSection was NOT expected to contain a TokenCode instance")
		return
	}

	if !code.IsSpecific() {
		t.Errorf("the TokenSection was expected to contain a SpecificTokenCode instance")
		return
	}

	if !code.Specific().HasAmount() {
		t.Errorf("the TokenSection was expected to contain a SpecificTokenCode instance with an amount")
		return
	}
}
