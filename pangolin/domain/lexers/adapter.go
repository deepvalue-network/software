package lexers

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

type adapter struct {
	builder Builder
	evts    []Event
	grammar grammar.Grammar
}

func createAdapter(builder Builder, evts []Event, grammar grammar.Grammar) Adapter {
	out := adapter{
		builder: builder,
		evts:    evts,
		grammar: grammar,
	}

	return &out
}

// ToLexer converts a script to a lexer
func (app *adapter) ToLexer(script string) (Lexer, error) {
	return app.builder.Create().WithGrammar(app.grammar).WithScript(string(script)).WithEvents(app.evts).Now()
}
