package parsers

import (
	"testing"
)

func Test_program_withApplication_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("program", grammarFile)

	file := "./tests/codes/program/application.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prog := ins.(Program)
	if !prog.IsApplication() {
		t.Errorf("the program was expected to be an Application")
		return
	}
}

func Test_program_withLanguage_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("program", grammarFile)

	file := "./tests/codes/program/language.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prog := ins.(Program)
	if !prog.IsLanguage() {
		t.Errorf("the program was expected to be a Language")
		return
	}
}

func Test_program_withScript_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("program", grammarFile)

	file := "./tests/codes/program/script.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prog := ins.(Program)
	if !prog.IsScript() {
		t.Errorf("the program was expected to be a Script")
		return
	}
}
