package rodan

import (
	"io/ioutil"
	"testing"

	"github.com/deepvalue-network/software/pangolin/bundles"
)

func TestPangolin_executeTests_Success(t *testing.T) {
	script, err := ioutil.ReadFile("./scripts/assignment/script.pangolin")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	currentDirPath := "./scripts/assignment"
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

	executable, err := pangolin.Linker().Execute(program)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	err = pangolin.Interpreter().Tests(executable)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

}
