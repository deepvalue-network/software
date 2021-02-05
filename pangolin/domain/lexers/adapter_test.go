package lexers

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

func TestConvert_toLexer_Success(t *testing.T) {
	// create the eventBuilder
	eventBuilder := NewEventBuilder()

	// create the event fn:
	eventFn := func(from int, to int, script []rune, rule grammar.Rule) []rune {
		script = append(script[:from], script[to:]...)
		return script
	}

	whiteSpaceEvent, err := eventBuilder.Create().WithToken("_whiteSpace").WithFn(eventFn).Now()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// create the events:
	events := []Event{
		whiteSpaceEvent,
	}

	grammarFilePath := "tests/grammar/grammar.json"
	adapter, err := NewAdapterBuilder().Create().WithGrammarFilePath(grammarFilePath).WithEvents(events).Now()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// read the code:
	script, err := ioutil.ReadFile("tests/code.rod")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// read the lexer:
	lexer, err := adapter.ToLexer(string(script))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if lexer == nil {
		t.Errorf("the lexer was expected to be valid, nil returned")
		return
	}

	log.Printf("\n %v \n", lexer)
}
