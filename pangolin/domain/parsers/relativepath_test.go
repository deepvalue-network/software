package parsers

import (
	"testing"
)

func Test_relativePath_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relativePath", grammarFile)

	file := "./tests/codes/relativepath/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	path := ins.(RelativePath)
	if len(path.All()) != 1 {
		t.Errorf("the path was expecting %d sections, %d returned", 1, len(path.All()))
		return
	}
}

func Test_relativePath_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relativePath", grammarFile)

	file := "./tests/codes/relativepath/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	path := ins.(RelativePath)
	if len(path.All()) != 4 {
		t.Errorf("the path was expecting %d sections, %d returned", 4, len(path.All()))
		return
	}
}
