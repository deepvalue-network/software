package parsers

import (
	"testing"
)

func Test_loadSingle_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("loadSingle", grammarFile)

	file := "./tests/codes/loadsingle/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	load := ins.(LoadSingle)
	if load.Internal() != "internal" {
		t.Errorf("the load internal variable was expected to be %s, %s returned", "internal", load.Internal())
		return
	}

	if load.External() != "external" {
		t.Errorf("the load external variable was expected to be %s, %s returned", "external", load.External())
		return
	}
}
