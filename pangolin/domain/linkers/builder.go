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
	programBuilder                  ProgramBuilder
	languageDefinitionBuilder       LanguageDefinitionBuilder
	pathsBuilder                    PathsBuilder
	scriptBuilder                   ScriptBuilder
	languageReferenceBuilder        LanguageReferenceBuilder
	languageApplicationBuilder      LanguageApplicationBuilder
	prevParser                      parsers.Parser
	dirPath                         string
}

func createBuilder(
	middleAdapter middle.Adapter,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	applicationBuilder ApplicationBuilder,
	languageBuilder LanguageBuilder,
	programBuilder ProgramBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	pathsBuilder PathsBuilder,
	scriptBuilder ScriptBuilder,
	languageReferenceBuilder LanguageReferenceBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
) Builder {
	out := builder{
		middleAdapter:                   nil,
		grammarRetrieverCriteriaBuilder: nil,
		applicationBuilder:              nil,
		languageBuilder:                 nil,
		programBuilder:                  nil,
		languageDefinitionBuilder:       nil,
		pathsBuilder:                    nil,
		scriptBuilder:                   nil,
		languageReferenceBuilder:        nil,
		languageApplicationBuilder:      nil,
		prevParser:                      nil,
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
		app.programBuilder,
		app.languageDefinitionBuilder,
		app.pathsBuilder,
		app.scriptBuilder,
		app.languageReferenceBuilder,
		app.languageApplicationBuilder,
	)
}

// WithPreviousParser adds a previous parser to the builder
func (app *builder) WithPreviousParser(prevParser parsers.Parser) Builder {
	app.prevParser = prevParser
	return app
}

// WithDirPath adds a directory path to the builder
func (app *builder) WithDirPath(dirPath string) Builder {
	app.dirPath = dirPath
	return app
}

// Now builds a new Linker instance
func (app *builder) Now() (Linker, error) {
	if app.prevParser == nil {
		return nil, errors.New("the previous parser is mandatory in order to build a Linker instance")
	}

	if app.dirPath == "" {
		return nil, errors.New("the directory path is mandatory in order to build a Linker instance")
	}

	return createLinker(
		app.middleAdapter,
		app.grammarRetrieverCriteriaBuilder,
		app.applicationBuilder,
		app.languageBuilder,
		app.programBuilder,
		app.languageDefinitionBuilder,
		app.pathsBuilder,
		app.scriptBuilder,
		app.languageReferenceBuilder,
		app.languageApplicationBuilder,
		app.prevParser,
		app.dirPath,
	), nil
}
