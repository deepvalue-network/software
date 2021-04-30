package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/compilers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

const scriptName = "default"

// NewCompiler creates a new compiler application
func NewCompiler(
	middleAdapter middle.Adapter,
	linker Linker,
	interpreterBuilder interpreters.Builder,
) Compiler {
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
	languageBuilder := linkers.NewLanguageBuilder()
	programBuilder := linkers.NewProgramBuilder()
	languageDefinitionBuilder := linkers.NewLanguageDefinitionBuilder()
	pathsBuilder := linkers.NewPathsBuilder()
	scriptBuilder := linkers.NewScriptBuilder()
	languageReferenceBuilder := linkers.NewLanguageReferenceBuilder()
	languageApplicationBuilder := linkers.NewLanguageApplicationBuilder()
	return createLinker(
		grammarRetrieverCriteriaBuilder,
		parserBuilder,
		parser,
		applicationBuilder,
		languageBuilder,
		programBuilder,
		languageDefinitionBuilder,
		pathsBuilder,
		scriptBuilder,
		languageReferenceBuilder,
		languageApplicationBuilder,
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
