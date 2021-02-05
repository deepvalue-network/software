package parsers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
)

type builder struct {
	eventsAdapter EventsAdapter
	lexer         lexers.Lexer
	params        []ToEventsParams
	replacements  map[string]RetrieveReplacementsFn
}

func createBuilder(eventsAdapter EventsAdapter) Builder {
	out := builder{
		eventsAdapter: eventsAdapter,
		lexer:         nil,
		params:        nil,
		replacements:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.eventsAdapter)
}

// WithLexer adds a lexer to the builder
func (app *builder) WithLexer(lexer lexers.Lexer) Builder {
	app.lexer = lexer
	return app
}

// WithEventParams add event params to the builder
func (app *builder) WithEventParams(params []ToEventsParams) Builder {
	app.params = params
	return app
}

// WithReplacement add replacements to the builder
func (app *builder) WithReplacements(replacements map[string]RetrieveReplacementsFn) Builder {
	app.replacements = replacements
	return app
}

// Now builds a new Parser instance
func (app *builder) Now() (Parser, error) {
	if app.lexer == nil {
		return nil, errors.New("the Lexer instance is mandatory in order to build a Parser instance")
	}

	if app.replacements != nil {
		for index, oneParam := range app.params {
			if repl, ok := app.replacements[oneParam.Token]; ok {
				app.params[index].RetrieveReplacement = repl
			}
		}
	}

	events, err := app.eventsAdapter.ToEvents(app.params)
	if err != nil {
		return nil, err
	}

	return createParser(app.lexer, events), nil
}
