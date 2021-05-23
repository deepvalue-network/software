package linkers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	middleAdapter                   middle.Adapter
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	applicationBuilder              ApplicationBuilder
	languageBuilder                 LanguageBuilder
	executableBuilder               ExecutableBuilder
	programBuilder                  ProgramBuilder
	languageDefinitionBuilder       LanguageDefinitionBuilder
	pathsBuilder                    PathsBuilder
	scriptBuilder                   ScriptBuilder
	testBuilder                     TestBuilder
	languageReferenceBuilder        LanguageReferenceBuilder
	languageApplicationBuilder      LanguageApplicationBuilder
	parser                          parsers.Parser
	dirPath                         string
}

func createBuilder(
	middleAdapter middle.Adapter,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	applicationBuilder ApplicationBuilder,
	languageBuilder LanguageBuilder,
	executableBuilder ExecutableBuilder,
	programBuilder ProgramBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	pathsBuilder PathsBuilder,
	scriptBuilder ScriptBuilder,
	testBuilder TestBuilder,
	languageReferenceBuilder LanguageReferenceBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
) Builder {
	out := builder{
		middleAdapter:                   middleAdapter,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		applicationBuilder:              applicationBuilder,
		languageBuilder:                 languageBuilder,
		executableBuilder:               executableBuilder,
		programBuilder:                  programBuilder,
		languageDefinitionBuilder:       languageDefinitionBuilder,
		pathsBuilder:                    pathsBuilder,
		scriptBuilder:                   scriptBuilder,
		testBuilder:                     testBuilder,
		languageReferenceBuilder:        languageReferenceBuilder,
		languageApplicationBuilder:      languageApplicationBuilder,
		parser:                          nil,
		dirPath:                         "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.middleAdapter,
		app.grammarRetrieverCriteriaBuilder,
		app.applicationBuilder,
		app.languageBuilder,
		app.executableBuilder,
		app.programBuilder,
		app.languageDefinitionBuilder,
		app.pathsBuilder,
		app.scriptBuilder,
		app.testBuilder,
		app.languageReferenceBuilder,
		app.languageApplicationBuilder,
	)
}

// WithParser adds a parser to the builder
func (app *builder) WithParser(parser parsers.Parser) Builder {
	app.parser = parser
	return app
}

// WithDirPath adds a directory path to the builder
func (app *builder) WithDirPath(dirPath string) Builder {
	app.dirPath = dirPath
	return app
}

// Now builds a new Linker instance
func (app *builder) Now() (Linker, error) {
	if app.parser == nil {
		return nil, errors.New("the parser is mandatory in order to build a Linker instance")
	}

	if app.dirPath == "" {
		return nil, errors.New("the directory path is mandatory in order to build a Linker instance")
	}

	return createLinker(
		app.middleAdapter,
		app.grammarRetrieverCriteriaBuilder,
		app.applicationBuilder,
		app.languageBuilder,
		app.executableBuilder,
		app.programBuilder,
		app.languageDefinitionBuilder,
		app.pathsBuilder,
		app.scriptBuilder,
		app.testBuilder,
		app.languageReferenceBuilder,
		app.languageApplicationBuilder,
		app.parser,
		app.dirPath,
	), nil
}
