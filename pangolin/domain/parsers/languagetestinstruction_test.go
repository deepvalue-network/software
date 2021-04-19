package parsers

import (
	"testing"
)

func Test_languageTestInstruction_withLanguage_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageTestInstruction", grammarFile)

	file := "./tests/codes/languagetestinstruction/language.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(LanguageTestInstruction)
	if !testIns.IsLanguageInstruction() {
		t.Errorf("the languageTestInstruction was expected to be a languageInstruction")
		return
	}

	if testIns.IsTestInstruction() {
		t.Errorf("the languageTestInstruction was NOT expected to be a testInstruction")
		return
	}
}

func Test_languageTestInstruction_withTest_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageTestInstruction", grammarFile)

	file := "./tests/codes/languagetestinstruction/test.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(LanguageTestInstruction)
	if testIns.IsLanguageInstruction() {
		t.Errorf("the languageTestInstruction was NOT expected to be a languageInstruction")
		return
	}

	if !testIns.IsTestInstruction() {
		t.Errorf("the languageTestInstruction was expected to be a testInstruction")
		return
	}
}
