package lexers

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

type scriptApplicationBuilder struct {
	grammar grammar.Grammar
	evts    []Event
}

func createScriptApplicationBuilder() ScriptApplicationBuilder {
	out := scriptApplicationBuilder{
		grammar: nil,
		evts:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptApplicationBuilder) Create() ScriptApplicationBuilder {
	return createScriptApplicationBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *scriptApplicationBuilder) WithGrammar(grammar grammar.Grammar) ScriptApplicationBuilder {
	app.grammar = grammar
	return app
}

// WithEvents add events to the builder
func (app *scriptApplicationBuilder) WithEvents(evts []Event) ScriptApplicationBuilder {
	app.evts = evts
	return app
}

// Now builds a new ScriptApplication instance
func (app *scriptApplicationBuilder) Now() (ScriptApplication, error) {
	if app.grammar == nil {
		return nil, errors.New("the Grammar is mandatory in order to build a ScriptApplication instance")
	}

	if app.evts == nil {
		app.evts = []Event{}
	}

	mp := map[string]Event{}
	for _, oneEvent := range app.evts {
		mp[oneEvent.Token()] = oneEvent
	}

	return createScriptApplication(app.grammar, mp), nil
}
