package parsers

import (
	"testing"
)

func Test_testable_withExecutable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testable", grammarFile)

	file := "./tests/codes/testable/executable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testable := ins.(Testable)
	if !testable.IsExecutable() {
		t.Errorf("the testable was expected to contain an Executable instance")
		return
	}

	if testable.IsLanguage() {
		t.Errorf("the executable was NOT expected to contain a LanguageApplication instance")
		return
	}
}

func Test_testable_withLanguageDefinition_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testable", grammarFile)

	file := "./tests/codes/testable/languagedefinition.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testable := ins.(Testable)
	if testable.IsExecutable() {
		t.Errorf("the testable was NOT expected to contain an Executable instance")
		return
	}

	if !testable.IsLanguage() {
		t.Errorf("the executable was expected to contain a LanguageDefinition instance")
		return
	}
}
