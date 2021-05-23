package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type languageInstructionCommonBuilder struct {
	insAppBuilder                   InstructionBuilder
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	lexerParserApplication          lexer_parser.Application
	lexerParserBuilder              lexer_parser.Builder
	lexerAdapterBuilder             lexers.AdapterBuilder
	langDef                         linkers.LanguageDefinition
	stackFrame                      stackframes.StackFrame
	state                           LanguageState
	labelFn                         CallLabelByNameFn
	events                          []lexers.Event
}

func createLanguageInstructionCommonBuilder(
	insAppBuilder InstructionBuilder,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerAdapterBuilder lexers.AdapterBuilder,
) LanguageInstructionCommonBuilder {
	out := languageInstructionCommonBuilder{
		insAppBuilder:                   insAppBuilder,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		lexerParserApplication:          lexerParserApplication,
		lexerParserBuilder:              lexerParserBuilder,
		lexerAdapterBuilder:             lexerAdapterBuilder,
		langDef:                         nil,
		stackFrame:                      nil,
		state:                           nil,
		labelFn:                         nil,
		events:                          nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageInstructionCommonBuilder) Create() LanguageInstructionCommonBuilder {
	return createLanguageInstructionCommonBuilder(
		app.insAppBuilder,
		app.grammarRetrieverCriteriaBuilder,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerAdapterBuilder,
	)
}

// WithLanguage adds a language to the builder
func (app *languageInstructionCommonBuilder) WithLanguage(langDef linkers.LanguageDefinition) LanguageInstructionCommonBuilder {
	app.langDef = langDef
	return app
}

// WithCallLabelFn adds a label callBack func to the builder
func (app *languageInstructionCommonBuilder) WithCallLabelFn(labelFn CallLabelByNameFn) LanguageInstructionCommonBuilder {
	app.labelFn = labelFn
	return app
}

// WithStackFrame adds a stackframe to the builder
func (app *languageInstructionCommonBuilder) WithStackFrame(stackFrame stackframes.StackFrame) LanguageInstructionCommonBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithState adds a state to the builder
func (app *languageInstructionCommonBuilder) WithState(state LanguageState) LanguageInstructionCommonBuilder {
	app.state = state
	return app
}

// WithEvents add events to the builder
func (app *languageInstructionCommonBuilder) WithEvents(events []lexers.Event) LanguageInstructionCommonBuilder {
	app.events = events
	return app
}

// Now builds a new LanguageInstructionCommon instance
func (app *languageInstructionCommonBuilder) Now() (LanguageInstructionCommon, error) {
	if app.langDef == nil {
		return nil, errors.New("the language definition is mandatory in order to build a LanguageInstructionCommon instance")
	}

	if app.labelFn == nil {
		return nil, errors.New("the labelCallFunc is mandatory in order to build a LanguageInstructionCommon instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a LanguageInstructionCommon instance")
	}

	if app.state == nil {
		return nil, errors.New("the state is mandatory in order to build a LanguageInstructionCommon instance")
	}

	root := app.langDef.Root()
	paths := app.langDef.Paths()

	patternMatches := map[string]definitions.PatternMatch{}
	patternMatchesList := app.langDef.PatternMatches()
	for _, onePatternMatch := range patternMatchesList {
		patternMatches[onePatternMatch.Pattern()] = onePatternMatch
	}

	baseDir := paths.BaseDir()
	tokensPath := paths.Tokens()
	rulesPath := paths.Rules()
	retrieverCriteriaBuilder := app.grammarRetrieverCriteriaBuilder.Create().
		WithRoot(root).
		WithBaseDirPath(baseDir).
		WithTokensPath(tokensPath).
		WithRulesPath(rulesPath)

	if paths.HasChannels() {
		channelsPath := paths.Channels()
		retrieverCriteriaBuilder.WithChannelsPath(channelsPath)
	}

	retrieverCriteria, err := retrieverCriteriaBuilder.Now()
	if err != nil {
		return nil, err
	}

	lexerAdapterBuilder := app.lexerAdapterBuilder.Create().WithGrammarRetrieverCriteria(retrieverCriteria).WithEvents(app.events)
	insApp, err := app.insAppBuilder.Create().WithCallLabelFn(app.labelFn).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	return createLanguageInstructionCommon(
		app.lexerParserApplication,
		app.lexerParserBuilder,
		lexerAdapterBuilder,
		patternMatches,
		insApp,
		app.state,
		app.stackFrame,
		app.labelFn,
	), nil
}
