package parsers

import (
	"testing"
)

func Test_languageInstructionCommon_withInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageInstructionCommon", grammarFile)

	file := "./tests/codes/languageinstructioncommon/instruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageInstructionCommon)
	if !langIns.IsInstruction() {
		t.Errorf("the LanguageInstruction was expected to contain an instruction")
		return
	}

	if langIns.IsMatch() {
		t.Errorf("the LanguageInstruction was NOT expected to contain a match")
		return
	}
}

func Test_languageInstructionCommon_withMatch_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageInstructionCommon", grammarFile)

	file := "./tests/codes/languageinstructioncommon/match.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageInstructionCommon)
	if langIns.IsInstruction() {
		t.Errorf("the LanguageInstruction was NOT expected to contain an instruction")
		return
	}

	if !langIns.IsMatch() {
		t.Errorf("the LanguageInstruction was expected to contain a match")
		return
	}
}
