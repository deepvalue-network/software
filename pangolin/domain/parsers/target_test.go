package parsers

import (
	"testing"
)

func Test_target_withOneEvent_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("target", grammarFile)

	file := "./tests/codes/target/one.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	target := ins.(Target)
	if target.Name() != "my_target" {
		t.Errorf("the name was expected to be %s, %s returned", "my_target", target.Name())
		return
	}

	evts := target.Events()
	amount := len(evts)
	if amount != 1 {
		t.Errorf("%d events were expected, %d returned", 1, amount)
		return
	}
}

func Test_target_withMultipleEvents_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("target", grammarFile)

	file := "./tests/codes/target/multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	target := ins.(Target)
	if target.Name() != "my_target" {
		t.Errorf("the name was expected to be %s, %s returned", "my_target", target.Name())
		return
	}

	evts := target.Events()
	amount := len(evts)
	if amount != 2 {
		t.Errorf("%d events were expected, %d returned", 2, amount)
		return
	}
}
