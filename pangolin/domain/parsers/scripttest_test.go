package parsers

import (
	"testing"
)

func Test_scriptTest_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptTest", grammarFile)

	file := "./tests/codes/scripttest/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	test := ins.(ScriptTest)
	if test.Name() != "testName" {
		t.Errorf("the name was expected to be %s, %s returned", "testName", test.Name())
		return
	}

	if test.Path().String() != "./../myDir/myFile.json" {
		t.Errorf("the path was expected to be %s, %s returned", "./../myDir/myFile.json", test.Path().String())
		return
	}
}
