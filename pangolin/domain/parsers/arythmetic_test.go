package parsers

import (
	"testing"
)

func Test_arythmetic_add_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("arythmetic", grammarFile)

	file := "./tests/codes/arythmetic/add.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ary := ins.(Arythmetic)
	if !ary.IsAdd() {
		t.Errorf("the Arythmetic was expected to contain an add operation")
		return
	}
}

func Test_arythmetic_sub_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("arythmetic", grammarFile)

	file := "./tests/codes/arythmetic/sub.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ary := ins.(Arythmetic)
	if !ary.IsSub() {
		t.Errorf("the Arythmetic was expected to contain a sub operation")
		return
	}
}

func Test_arythmetic_mul_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("arythmetic", grammarFile)

	file := "./tests/codes/arythmetic/mul.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ary := ins.(Arythmetic)
	if !ary.IsMul() {
		t.Errorf("the Arythmetic was expected to contain a mul operation")
		return
	}
}

func Test_arythmetic_div_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("arythmetic", grammarFile)

	file := "./tests/codes/arythmetic/div.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ary := ins.(Arythmetic)
	if !ary.IsDiv() {
		t.Errorf("the Arythmetic was expected to contain a div operation")
		return
	}
}
