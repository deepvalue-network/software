package parsers

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
)

func createParserForTests(rootPattern string, grammarFile string) Parser {
	wsEvent, err := NewWhiteSpaceEvent("_whiteSpace")
	if err != nil {
		panic(err)
	}

	lexerAdapter, err := lexers.NewAdapterBuilder().WithRoot(rootPattern).WithGrammarFilePath(grammarFile).WithEvents([]lexers.Event{
		wsEvent,
	}).Now()

	if err != nil {
		panic(err)
	}

	parser, err := NewParserBuilder().Create().WithLexerAdapter(lexerAdapter).Now()
	if err != nil {
		panic(err)
	}

	return parser
}
