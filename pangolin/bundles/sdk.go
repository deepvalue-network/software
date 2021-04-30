package bundles

import (
	"github.com/deepvalue-network/software/pangolin/application"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewInterpreter creates a new interpreter instance
func NewInterpreter(dirPath string, langPath string, grammarFilePath string, name string) (interpreters.Interpreter, error) {
	events, err := NewEvents()
	if err != nil {
		return nil, err
	}

	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	machineLanguageBuilder := interpreters.NewMachineLanguageBuilder(lexerAdapterBuilder, events)
	interpreterBuilder := interpreters.NewBuilder(machineLanguageBuilder)
	parserLinker, err := NewParserLinker(dirPath, grammarFilePath, name)
	if err != nil {
		return nil, err
	}

	program, err := parserLinker.File(langPath)
	if err != nil {
		return nil, err
	}

	return interpreterBuilder.Create().WithProgram(program).Now()
}

// NewParserLinker creates a new parser linker application
func NewParserLinker(dirPath string, grammarFilePath string, name string) (application.ParserLinker, error) {
	middleAdapter := middle.NewAdapter()
	grammarRetrieverCriteriaRepository, err := grammar.NewRetrieverCriteriaRepositoryBuilder().WithName(name).Now()
	if err != nil {
		return nil, err
	}

	grammarRetrieverCriteria, err := grammarRetrieverCriteriaRepository.Retrieve(grammarFilePath)
	if err != nil {
		return nil, err
	}

	// application parser:
	appParserBuilder := application.NewParserBuilder(middleAdapter, "_whiteSpace")
	appParser, err := appParserBuilder.Create().WithRetrieverCriteria(grammarRetrieverCriteria).Now()
	if err != nil {
		return nil, err
	}

	// application linker:
	appLinker := application.NewLinker(
		appParserBuilder,
		appParser,
		dirPath,
	)

	// application parser linker:
	return application.NewParserLinker(appParser, appLinker), nil
}

// NewParser creates a new parser instance
func NewParser(grammarFilePath string, name string) (parsers.Parser, error) {
	events, err := NewEvents()
	if err != nil {
		return nil, err
	}

	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	lexerAdapter, err := lexerAdapterBuilder.Create().WithGrammarFilePath(grammarFilePath).WithEvents(events).Now()
	if err != nil {
		return nil, err
	}

	parserBuilder := parsers.NewParserBuilder()
	return parserBuilder.Create().WithLexerAdapter(lexerAdapter).Now()
}

// NewEvents creates new events instance
func NewEvents() ([]lexers.Event, error) {
	wsEvent, err := parsers.NewWhiteSpaceEvent("_whiteSpace")
	if err != nil {
		return nil, err
	}

	return []lexers.Event{
		wsEvent,
	}, nil
}
