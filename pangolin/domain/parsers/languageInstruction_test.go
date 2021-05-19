package parsers

import (
	"testing"
)

func Test_languageInstruction_withInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageInstruction", grammarFile)

	file := "./tests/codes/languageinstruction/instruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageInstruction)
	if !langIns.IsInstruction() {
		t.Errorf("the LanguageInstruction was expected to contain an instruction")
		return
	}

	if langIns.IsCommand() {
		t.Errorf("the LanguageInstruction was NOT expected to contain a command")
		return
	}
}

func Test_languageInstruction_withCommand_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageInstruction", grammarFile)

	file := "./tests/codes/languageinstruction/command.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	langIns := ins.(LanguageInstruction)
	if langIns.IsInstruction() {
		t.Errorf("the LanguageInstruction was NOT expected to contain an instruction")
		return
	}

	if !langIns.IsCommand() {
		t.Errorf("the LanguageInstruction was expected to contain a command")
		return
	}
}
