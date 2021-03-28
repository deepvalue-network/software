package parsers

import (
	"testing"
)

func Test_jump_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("jump", grammarFile)

	file := "./tests/codes/jump/jump.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	jmp := ins.(Jump)
	if jmp.HasCondition() {
		t.Errorf("the Jump was not expected to contain a condition")
		return
	}

	lbl := jmp.Label()
	if lbl != "myLabel" {
		t.Errorf("the Label was expected to be %s, %s returned", "myLabel", lbl)
		return
	}
}

func Test_jumpIf_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("jump", grammarFile)

	file := "./tests/codes/jump/jump_if.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	jmp := ins.(Jump)
	if !jmp.HasCondition() {
		t.Errorf("the Jump was expected to contain a condition")
		return
	}

	lbl := jmp.Label()
	if lbl != "myLabel" {
		t.Errorf("the Label was expected to be %s, %s returned", "myLabel", lbl)
		return
	}
}
