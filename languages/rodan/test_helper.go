package rodan

import (
	scripts "github.com/steve-care-software/products/pangolin/application"
	"github.com/steve-care-software/products/pangolin/domain/interpreters"
	"github.com/steve-care-software/products/pangolin/domain/lexers"
	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
	"github.com/steve-care-software/products/pangolin/domain/middle"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

func create(dirPath string) (scripts.ParserLinker, interpreters.Builder, error) {
	name := "pangolin"
	grammarFile := "../../pangolin/domain/parsers/grammar/grammar.json"
	wsEvent, err := parsers.NewWhiteSpaceEvent("_whiteSpace")
	if err != nil {
		return nil, nil, err
	}

	events := []lexers.Event{
		wsEvent,
	}

	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	lexerAdapter, err := lexerAdapterBuilder.Create().WithGrammarFilePath(grammarFile).WithEvents(events).Now()
	if err != nil {
		return nil, nil, err
	}

	parserBuilder := parsers.NewParserBuilder()
	parser, err := parserBuilder.Create().WithLexerAdapter(lexerAdapter).Now()
	if err != nil {
		return nil, nil, err
	}

	middleAdapter := middle.NewAdapter(parser)
	grammarRetrieverCriteriaRepository, err := grammar.NewRetrieverCriteriaRepositoryBuilder().WithName(name).Now()
	if err != nil {
		return nil, nil, err
	}

	grammarRetrieverCriteria, err := grammarRetrieverCriteriaRepository.Retrieve(grammarFile)
	if err != nil {
		return nil, nil, err
	}

	// application parser:
	appParserBuilder := scripts.NewParserBuilder(middleAdapter, "_whiteSpace")
	appParser, err := appParserBuilder.Create().WithRetrieverCriteria(grammarRetrieverCriteria).Now()
	if err != nil {
		return nil, nil, err
	}

	// application linker:
	appLinker := scripts.NewLinker(
		appParserBuilder,
		appParser,
		dirPath,
	)

	// application parser linker:
	appParserLinker := scripts.NewParserLinker(appParser, appLinker)

	// interpreter:
	machineBuilder := interpreters.NewMachineBuilder(lexerAdapterBuilder, events)
	interpreterBuilder := interpreters.NewBuilder(machineBuilder)

	// returns:
	return appParserLinker, interpreterBuilder, nil
}
