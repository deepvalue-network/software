package application

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	lexerAdapterBuilder lexers.AdapterBuilder
	parserBuilder       parsers.ParserBuilder
	linkerBuilder       linkers.Builder
	interpreterBuilder  interpreters.Builder
	dirPath             string
	grammarFilePath     string
	events              []lexers.Event
}

func createBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
	parserBuilder parsers.ParserBuilder,
	linkerBuilder linkers.Builder,
	interpreterBuilder interpreters.Builder,
) Builder {
	out := builder{
		lexerAdapterBuilder: lexerAdapterBuilder,
		parserBuilder:       parserBuilder,
		linkerBuilder:       linkerBuilder,
		interpreterBuilder:  interpreterBuilder,
		dirPath:             "",
		grammarFilePath:     "",
		events:              nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.lexerAdapterBuilder,
		app.parserBuilder,
		app.linkerBuilder,
		app.interpreterBuilder,
	)
}

// WithCurrentDirPath adds the current dir path to the builder
func (app *builder) WithCurrentDirPath(dirPath string) Builder {
	app.dirPath = dirPath
	return app
}

// WithGrammarFilePath adds a grammar file path to the builder
func (app *builder) WithGrammarFilePath(grammarFilePath string) Builder {
	app.grammarFilePath = grammarFilePath
	return app
}

// WithEvents adds events to the builder
func (app *builder) WithEvents(events []lexers.Event) Builder {
	app.events = events
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.dirPath == "" {
		return nil, errors.New("the current directory path is mandatory in order to build an Application instance")
	}

	lexerAdapter, err := app.lexerAdapterBuilder.Create().WithGrammarFilePath(app.grammarFilePath).WithEvents(app.events).Now()
	if err != nil {
		return nil, err
	}

	parser, err := app.parserBuilder.Create().WithLexerAdapter(lexerAdapter).Now()
	if err != nil {
		return nil, err
	}

	linker, err := app.linkerBuilder.Create().WithParser(parser).WithDirPath(app.dirPath).Now()
	if err != nil {
		return nil, err
	}

	interpreter, err := app.interpreterBuilder.Create().WithEvents(app.events).WithLinker(linker).WithParser(parser).Now()
	if err != nil {
		return nil, err
	}

	lexerApp := createLexer(lexerAdapter)
	parserApp := createParser(parser)
	linkerApp := createLinker(linker)
	interpreterApp := createInterpreter(interpreter)
	return createApplication(lexerApp, parserApp, linkerApp, interpreterApp), nil
}
