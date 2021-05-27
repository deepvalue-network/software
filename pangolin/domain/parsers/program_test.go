package parsers

import (
	"testing"
)

func Test_program_withTestable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("program", grammarFile)

	file := "./tests/codes/program/testable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prog := ins.(Program)
	if !prog.IsTestable() {
		t.Errorf("the program was expected to be a Testable instance")
		return
	}

	if prog.IsLanguage() {
		t.Errorf("the program was NOT expected to be a LanguageApplication instance")
		return
	}
}

func Test_program_withLanguageApplication_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("program", grammarFile)

	file := "./tests/codes/program/languageapplication.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prog := ins.(Program)
	if prog.IsTestable() {
		t.Errorf("the program was NOT expected to be a Testable instance")
		return
	}

	if !prog.IsLanguage() {
		t.Errorf("the program was expected to be a LanguageApplication instance")
		return
	}
}
