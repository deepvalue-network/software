package parsers

import (
	"testing"
)

func Test_labelDeclaration_single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelDeclaration", grammarFile)

	file := "./tests/codes/labeldeclaration/single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(LabelDeclaration)
	if decl.Name() != "myLabel" {
		t.Errorf("the LabelDeclaration was expected '%s' as a name, '%s' returned", "myLabel", decl.Name())
		return
	}

	lst := decl.Instructions()
	if len(lst) != 1 {
		t.Errorf("the LabelDeclaration was expected to contain %d instructions, %d returned", 1, len(lst))
		return
	}
}

func Test_labelDeclaration_multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelDeclaration", grammarFile)

	file := "./tests/codes/labeldeclaration/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	decl := ins.(LabelDeclaration)
	if decl.Name() != "myLabelAgain" {
		t.Errorf("the LabelDeclaration was expected '%s' as a name, '%s' returned", "myLabelAgain", decl.Name())
		return
	}

	lst := decl.Instructions()
	if len(lst) != 3 {
		t.Errorf("the LabelDeclaration was expected to contain %d instructions, %d returned", 3, len(lst))
		return
	}
}
