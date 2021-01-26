package parsers

import (
	"testing"
)

func Test_importSingle_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("importSingle", grammarFile)

	file := "./tests/codes/importsingle/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	imp := ins.(ImportSingle)
	if imp.Name() != "internal_name" {
		t.Errorf("the name was expected to be %s, %s returned", "internal_name", imp.Name())
		return
	}

	if imp.Path() == nil {
		t.Errorf("the path was expected to be valid")
		return
	}
}
