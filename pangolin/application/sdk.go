package scripts

import (
	"github.com/steve-care-software/products/pangolin/domain/compilers"
	"github.com/steve-care-software/products/pangolin/domain/interpreters"
	"github.com/steve-care-software/products/pangolin/domain/lexers"
	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
	"github.com/steve-care-software/products/pangolin/domain/linkers"
	"github.com/steve-care-software/products/pangolin/domain/middle"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

const scriptName = "default"

// NewCompiler creates a new compiler application
func NewCompiler(
	parser parsers.Parser,
	linker Linker,
	interpreterBuilder interpreters.Builder,
) Compiler {
	middleAdapter := middle.NewAdapter(parser)
	compilerApplication := compilers.NewApplication(middleAdapter, interpreterBuilder)
	return createCompiler(compilerApplication, linker)
}

// NewParserLinker creates a new parserLinker application
func NewParserLinker(
	parser Parser,
	linker Linker,
) ParserLinker {
	return createParserLinker(parser, linker)
}

// NewLinker creates a new linker instance
func NewLinker(
	parserBuilder ParserBuilder,
	parser Parser,
	dirPath string,
) Linker {
	grammarRetrieverCriteriaBuilder := grammar.NewRetrieverCriteriaBuilder()
	applicationBuilder := linkers.NewApplicationBuilder()
	programBuilder := linkers.NewProgramBuilder()
	languageBuilder := linkers.NewLanguageBuilder()
	pathsBuilder := linkers.NewPathsBuilder()
	scriptBuilder := linkers.NewScriptBuilder()
	languageReferenceBuilder := linkers.NewLanguageReferenceBuilder()
	return createLinker(
		grammarRetrieverCriteriaBuilder,
		parserBuilder,
		parser,
		applicationBuilder,
		programBuilder,
		languageBuilder,
		pathsBuilder,
		scriptBuilder,
		languageReferenceBuilder,
		dirPath,
	)
}

//NewParserBuilder creates a new ParserBuilder instance
func NewParserBuilder(middleAdapter middle.Adapter, whiteSpaceChannelName string) ParserBuilder {
	wsEvent, err := parsers.NewWhiteSpaceEvent(whiteSpaceChannelName)
	if err != nil {
		panic(err)
	}

	events := []lexers.Event{
		wsEvent,
	}

	pBuilder := parsers.NewParserBuilder()
	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	grammarRetrieverCriteriaBuilder := grammar.NewRetrieverCriteriaBuilder()
	return createParserBuilder(middleAdapter, pBuilder, lexerAdapterBuilder, grammarRetrieverCriteriaBuilder, events)
}

// ParserLinker represents a parser linker application
type ParserLinker interface {
	File(filePath string) (linkers.Program, error)
	Script(script string) (linkers.Program, error)
}

// Compiler represents a compiler application
type Compiler interface {
	Execute(script linkers.Script) (linkers.Application, error)
}

// Linker represents a linker application
type Linker interface {
	Execute(program middle.Program) (linkers.Program, error)
}

// ParserBuilder represents a parser builder
type ParserBuilder interface {
	Create() ParserBuilder
	WithRetrieverCriteria(criteria grammar.RetrieverCriteria) ParserBuilder
	Now() (Parser, error)
}

// Parser represents a parser
type Parser interface {
	File(filePath string) (middle.Program, error)
	Script(script string) (middle.Program, error)
}
