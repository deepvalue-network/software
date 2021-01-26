package parsers

import (
	"testing"
)

func Test_patternLabels_withEnter_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("patternLabels", grammarFile)

	file := "./tests/codes/patternlabels/with_enter.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lbls := ins.(PatternLabels)
	if !lbls.HasEnterLabel() {
		t.Errorf("the enter label was expected to be valid")
		return
	}

	if lbls.EnterLabel() != "myEnter" {
		t.Errorf("the enter label was expected to be %s, %s returned", "myEnter", lbls.EnterLabel())
		return
	}

	if lbls.HasExitLabel() {
		t.Errorf("the exit label was expected to be invalid")
		return
	}
}

func Test_patternLabels_withExit_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("patternLabels", grammarFile)

	file := "./tests/codes/patternlabels/with_exit.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lbls := ins.(PatternLabels)
	if lbls.HasEnterLabel() {
		t.Errorf("the enter label was expected to be invalid")
		return
	}

	if !lbls.HasExitLabel() {
		t.Errorf("the exit label was expected to be valid")
		return
	}

	if lbls.ExitLabel() != "myExit" {
		t.Errorf("the exit label was expected to be %s, %s returned", "myExit", lbls.ExitLabel())
		return
	}
}

func Test_patternLabels_withEnter_withExit_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("patternLabels", grammarFile)

	file := "./tests/codes/patternlabels/with_enter_with_exit.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lbls := ins.(PatternLabels)
	if !lbls.HasEnterLabel() {
		t.Errorf("the enter label was expected to be valid")
		return
	}

	if lbls.EnterLabel() != "myEnter" {
		t.Errorf("the enter label was expected to be %s, %s returned", "myEnter", lbls.EnterLabel())
		return
	}

	if !lbls.HasExitLabel() {
		t.Errorf("the exit label was expected to be valid")
		return
	}

	if lbls.ExitLabel() != "myExit" {
		t.Errorf("the exit label was expected to be %s, %s returned", "myExit", lbls.ExitLabel())
		return
	}
}
