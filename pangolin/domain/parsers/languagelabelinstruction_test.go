package parsers

import (
	"testing"
)

func Test_languageLabelInstruction_withLabelInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageLabelInstruction", grammarFile)

	file := "./tests/codes/languagelabelinstruction/labelinstruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageLabelInstruction)
	if !langIns.IsLabelInstruction() {
		t.Errorf("the LanguageLabelInstruction was expected to contain a label instruction")
		return
	}

	if langIns.IsLanguageInstruction() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a language instruction")
		return
	}

	if langIns.IsToken() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a token")
		return
	}
}

func Test_languageLabelInstruction_withLanguageInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageLabelInstruction", grammarFile)

	file := "./tests/codes/languagelabelinstruction/languageinstruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageLabelInstruction)
	if langIns.IsLabelInstruction() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a label instruction")
		return
	}

	if !langIns.IsLanguageInstruction() {
		t.Errorf("the LanguageLabelInstruction was expected to contain a language instruction")
		return
	}

	if langIns.IsToken() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a token")
		return
	}
}

func Test_languageLabelInstruction_withToken_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageLabelInstruction", grammarFile)

	file := "./tests/codes/languagelabelinstruction/token.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageLabelInstruction)
	if langIns.IsLabelInstruction() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a label instruction")
		return
	}

	if langIns.IsLanguageInstruction() {
		t.Errorf("the LanguageLabelInstruction was NOT expected to contain a language instruction")
		return
	}

	if !langIns.IsToken() {
		t.Errorf("the LanguageLabelInstruction was expected to contain a token")
		return
	}
}
