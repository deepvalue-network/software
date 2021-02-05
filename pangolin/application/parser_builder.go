package scripts

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type parserBuilder struct {
	middleAdapter                   middle.Adapter
	parserBuilder                   parsers.ParserBuilder
	lexerAdapterBuilder             lexers.AdapterBuilder
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	events                          []lexers.Event
	criteria                        grammar.RetrieverCriteria
}

func createParserBuilder(
	middleAdapter middle.Adapter,
	pBuilder parsers.ParserBuilder,
	lexerAdapterBuilder lexers.AdapterBuilder,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	events []lexers.Event,
) ParserBuilder {
	out := parserBuilder{
		middleAdapter:                   middleAdapter,
		parserBuilder:                   pBuilder,
		lexerAdapterBuilder:             lexerAdapterBuilder,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		events:                          events,
		criteria:                        nil,
	}

	return &out
}

// Create initializes the builder
func (app *parserBuilder) Create() ParserBuilder {
	return createParserBuilder(
		app.middleAdapter,
		app.parserBuilder,
		app.lexerAdapterBuilder,
		app.grammarRetrieverCriteriaBuilder,
		app.events,
	)
}

// WithRetrieverCriteria adds a grammar retrieverCriteria to the builder
func (app *parserBuilder) WithRetrieverCriteria(criteria grammar.RetrieverCriteria) ParserBuilder {
	app.criteria = criteria
	return app
}

// Now builds a new Parser instance
func (app *parserBuilder) Now() (Parser, error) {
	if app.criteria == nil {
		return nil, errors.New("the grammar's RetrieverCriteria is mandatory in order to build a Parser intance")
	}

	lexerAdapter, err := app.lexerAdapterBuilder.Create().
		WithGrammarRetrieverCriteria(app.criteria).
		WithEvents(app.events).
		Now()

	if err != nil {
		return nil, err
	}

	parser, err := app.parserBuilder.Create().WithLexerAdapter(lexerAdapter).Now()
	if err != nil {
		return nil, err
	}

	return createParser(app.middleAdapter, parser), nil
}
