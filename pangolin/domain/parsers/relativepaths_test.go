package parsers

import (
	"testing"
)

func Test_relativePaths_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relativePaths", grammarFile)

	file := "./tests/codes/relativepaths/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	relPaths := ins.(RelativePaths)
	paths := relPaths.All()
	if len(paths) != 1 {
		t.Errorf("%d RelativePath were expected, %d returned", 1, len(paths))
		return
	}
}

func Test_relativePaths_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("relativePaths", grammarFile)

	file := "./tests/codes/relativepaths/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	relPaths := ins.(RelativePaths)
	paths := relPaths.All()
	if len(paths) != 3 {
		t.Errorf("%d RelativePath were expected, %d returned", 3, len(paths))
		return
	}
}
