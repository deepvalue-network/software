package parsers

import (
	"testing"
)

func Test_event_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("event", grammarFile)

	file := "./tests/codes/event/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	evt := ins.(Event)
	if evt.Name() != "my_event" {
		t.Errorf("the name was expected to be %s, %s returned", "my_event", evt.Name())
		return
	}

	if evt.Label() != "MyLabel" {
		t.Errorf("the name was expected to be %s, %s returned", "MyLabel", evt.Label())
		return
	}
}
