package parsers

import (
	"testing"
)

func Test_token_withCodeMatch_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("token", grammarFile)

	file := "./tests/codes/token/codematch.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	token := ins.(Token)
	if !token.IsCodeMatch() {
		t.Errorf("the Token was expected to contain a CodeMatch")
		return
	}

	if token.IsTokenSection() {
		t.Errorf("the Token was NOT expected to contain a TokenSection")
		return
	}
}

func Test_token_withTokenSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("token", grammarFile)

	file := "./tests/codes/token/tokensection.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	token := ins.(Token)
	if token.IsCodeMatch() {
		t.Errorf("the Token was NOT expected to contain a CodeMatch")
		return
	}

	if !token.IsTokenSection() {
		t.Errorf("the Token was expected to contain a TokenSection")
		return
	}
}
