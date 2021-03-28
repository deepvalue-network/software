package parsers

import (
	"testing"
)

func Test_trigger_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("trigger", grammarFile)

	file := "./tests/codes/trigger/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	trigger := ins.(Trigger)
	if trigger.Variable() != "myStackVariableName" {
		t.Errorf("the variableName was expected to be %s, %s returned", "myStackVariableName", trigger.Variable())
		return
	}

	if trigger.Event() != "my_event" {
		t.Errorf("the event was expected to be %s, %s returned", "my_event", trigger.Event())
		return
	}
}
