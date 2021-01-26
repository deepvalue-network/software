package parsers

import (
	"testing"
)

func Test_testInstruction_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testDeclaration", grammarFile)

	file := "./tests/codes/testdeclaration/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(TestDeclaration)
	if decl.Name() != "myTest" {
		t.Errorf("the name was expected to be %s, %s returned", "myTest", decl.Name())
		return
	}

	lst := decl.Instructions()
	if len(lst) != 1 {
		t.Errorf("the testInstruction was expecting %d instructions, %d returned", 1, len(lst))
		return
	}
}

func Test_testInstruction_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testDeclaration", grammarFile)

	file := "./tests/codes/testdeclaration/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(TestDeclaration)
	if decl.Name() != "myTest" {
		t.Errorf("the name was expected to be %s, %s returned", "myTest", decl.Name())
		return
	}

	lst := decl.Instructions()
	if len(lst) != 2 {
		t.Errorf("the testInstruction was expecting %d instructions, %d returned", 2, len(lst))
		return
	}
}
