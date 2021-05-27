package rodan

import (
	"io/ioutil"
	"testing"

	"github.com/deepvalue-network/software/pangolin/bundles"
)

func TestPangolin_executeTests_Success(t *testing.T) {
	script, err := ioutil.ReadFile("./language.pangolin")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	currentDirPath := "./"
	grammarFile := "../../pangolin/domain/parsers/grammar/grammar.json"
	pangolin := bundles.NewPangolin(grammarFile, currentDirPath)
	lexer, err := pangolin.Lexer().Execute(string(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	program, err := pangolin.Parser().Execute(lexer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	linkedProgram, err := pangolin.Linker().Execute(program)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	if !linkedProgram.IsTestable() {
		t.Errorf("the linked program was expected to be a testable instance")
		return
	}

	linkedTestable := linkedProgram.Testable()
	err = pangolin.Interpreter().Tests(linkedTestable)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

}
