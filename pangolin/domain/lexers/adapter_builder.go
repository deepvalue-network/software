package lexers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

type adapterBuilder struct {
	builder                  Builder
	grammarBuilder           grammar.Builder
	grammarRepositoryBuilder grammar.RepositoryBuilder
	defaultName              string
	name                     string
	root                     string
	fileFetcher              grammar.FileFetcher
	grammarFilePath          string
	grammarRetrieverCriteria grammar.RetrieverCriteria
	grammar                  grammar.Grammar
	events                   []Event
}

func createAdapterBuilder(
	builder Builder,
	grammarBuilder grammar.Builder,
	grammarRepositoryBuilder grammar.RepositoryBuilder,
	defaultName string,
) AdapterBuilder {
	out := adapterBuilder{
		builder:                  builder,
		grammarBuilder:           grammarBuilder,
		grammarRepositoryBuilder: grammarRepositoryBuilder,
		defaultName:              defaultName,
		name:                     "",
		root:                     "",
		fileFetcher:              nil,
		grammarFilePath:          "",
		grammarRetrieverCriteria: nil,
		grammar:                  nil,
		events:                   nil,
	}

	return &out
}

// Create initialies the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(app.builder, app.grammarBuilder, app.grammarRepositoryBuilder, app.defaultName)
}

// WithName adds a name to the builder
func (app *adapterBuilder) WithName(name string) AdapterBuilder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *adapterBuilder) WithRoot(root string) AdapterBuilder {
	app.root = root
	return app
}

// WithFileFetcher adds a fileFetcher to the builder
func (app *adapterBuilder) WithFileFetcher(fileFetcher grammar.FileFetcher) AdapterBuilder {
	app.fileFetcher = fileFetcher
	return app
}

// WithGrammarFilePath adds a grammar filePath to the builder
func (app *adapterBuilder) WithGrammarFilePath(grammarFilePath string) AdapterBuilder {
	app.grammarFilePath = grammarFilePath
	return app
}

// WithGrammarRetrieverCriteria adds a grammar retrieverCriteria to the builder
func (app *adapterBuilder) WithGrammarRetrieverCriteria(grammarRetrieverCriteria grammar.RetrieverCriteria) AdapterBuilder {
	app.grammarRetrieverCriteria = grammarRetrieverCriteria
	return app
}

// WithGrammar adds a grammar to the builder
func (app *adapterBuilder) WithGrammar(grammar grammar.Grammar) AdapterBuilder {
	app.grammar = grammar
	return app
}

// WithEvents add events to the builder
func (app *adapterBuilder) WithEvents(events []Event) AdapterBuilder {
	app.events = events
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.name == "" {
		app.name = app.defaultName
	}

	grammarRepositoryBuilder := app.grammarRepositoryBuilder.Create()
	if app.fileFetcher != nil {
		grammarRepositoryBuilder.WithFileFetcher(app.fileFetcher)
	}

	grammarRepository, err := grammarRepositoryBuilder.Now()
	if err != nil {
		return nil, err
	}

	if app.grammarFilePath != "" {
		grammar, err := grammarRepository.RetrieveFromFile(app.root, app.name, app.grammarFilePath)
		if err != nil {
			return nil, err
		}

		app.grammar = grammar
	}

	if app.grammarRetrieverCriteria != nil {
		grammar, err := grammarRepository.Retrieve(app.grammarRetrieverCriteria)
		if err != nil {
			return nil, err
		}

		app.grammar = grammar
	}

	if app.root != "" {
		name := app.grammar.Name()

		mpRules := app.grammar.Rules()
		rules := []grammar.Rule{}
		for _, oneRule := range mpRules {
			rules = append(rules, oneRule)
		}

		builder := app.grammarBuilder.Create().WithName(name).WithRoot(app.root).WithRules(rules)
		if app.grammar.HasChannels() {
			channels := app.grammar.Channels()
			builder.WithChannels(channels)
		}

		if app.grammar.HasTokens() {
			tokens := app.grammar.Tokens()
			builder.WithTokens(tokens)
		}

		if app.grammar.HasSubGrammars() {
			sub := app.grammar.SubGrammars()
			builder.WithSubGrammars(sub)
		}

		grammar, err := builder.Now()
		if err != nil {
			return nil, err
		}

		app.grammar = grammar
	}

	if app.grammar == nil {
		return nil, errors.New("the grammarFilePath, the grammarRetrieverCriteria or a grammar instance was expected in order to build a lexer adapter instance")
	}

	if app.events == nil {
		app.events = []Event{}
	}

	return createAdapter(app.builder, app.events, app.grammar), nil
}
