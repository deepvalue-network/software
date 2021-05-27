package parsers

import (
	"testing"
)

func Test_executable_withApplication_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("executable", grammarFile)

	file := "./tests/codes/executable/application.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	executable := ins.(Executable)
	if !executable.IsApplication() {
		t.Errorf("the executable was expected to contain an Application instance")
		return
	}

	if executable.IsScript() {
		t.Errorf("the executable was NOT expected to contain a Script instance")
		return
	}
}

func Test_executable_withScript_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("executable", grammarFile)

	file := "./tests/codes/executable/script.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	executable := ins.(Executable)
	if executable.IsApplication() {
		t.Errorf("the executable was NOT expected to contain an Application instance")
		return
	}

	if !executable.IsScript() {
		t.Errorf("the executable was expected to contain a Script instance")
		return
	}
}
