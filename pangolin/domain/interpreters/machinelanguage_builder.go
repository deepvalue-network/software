package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type machineLanguageBuilder struct {
	variableBuilder                 var_variable.Builder
	valueBuilder                    var_value.Builder
	computableValueBuilder          computable.Builder
	lexerParserApplication          lexer_parser.Application
	lexerParserBuilder              lexer_parser.Builder
	lexerAdapterBuilder             lexers.AdapterBuilder
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	stackFrameBuilder               StackFrameBuilder
	events                          []lexers.Event
	lang                            linkers.LanguageDefinition
	input                           map[string]computable.Value
	fetchStackFrameFn               FetchStackFrameFunc
	machine                         Machine
}

func createMachineLanguageBuilder(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerAdapterBuilder lexers.AdapterBuilder,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	stackFrameBuilder StackFrameBuilder,
	events []lexers.Event,
) MachineLanguageBuilder {
	out := machineLanguageBuilder{
		variableBuilder:                 variableBuilder,
		valueBuilder:                    valueBuilder,
		computableValueBuilder:          computableValueBuilder,
		lexerParserApplication:          lexerParserApplication,
		lexerParserBuilder:              lexerParserBuilder,
		lexerAdapterBuilder:             lexerAdapterBuilder,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		stackFrameBuilder:               stackFrameBuilder,
		events:                          events,
		lang:                            nil,
		input:                           nil,
		fetchStackFrameFn:               nil,
		machine:                         nil,
	}

	return &out
}

// Create initializes the builder
func (app *machineLanguageBuilder) Create() MachineLanguageBuilder {
	return createMachineLanguageBuilder(
		app.variableBuilder,
		app.valueBuilder,
		app.computableValueBuilder,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerAdapterBuilder,
		app.grammarRetrieverCriteriaBuilder,
		app.stackFrameBuilder,
		app.events,
	)
}

// WithLanguage adds a language definition to the builder
func (app *machineLanguageBuilder) WithLanguage(lang linkers.LanguageDefinition) MachineLanguageBuilder {
	app.lang = lang
	return app
}

// WithInput adds input values to the builder
func (app *machineLanguageBuilder) WithInput(input map[string]computable.Value) MachineLanguageBuilder {
	app.input = input
	return app
}

// WithFetchStackFunc adds a fetchStackFunc func to the builder
func (app *machineLanguageBuilder) WithFetchStackFunc(fetchStackFrameFn FetchStackFrameFunc) MachineLanguageBuilder {
	app.fetchStackFrameFn = fetchStackFrameFn
	return app
}

// WithMachine adds a machine to the builder
func (app *machineLanguageBuilder) WithMachine(machine Machine) MachineLanguageBuilder {
	app.machine = machine
	return app
}

// Now builds a new Machine instance
func (app *machineLanguageBuilder) Now() (MachineLanguage, error) {
	if app.lang == nil {
		return nil, errors.New("the language definition is mandatory in order to build a MachineLanguage instance")
	}

	if app.fetchStackFrameFn == nil {
		return nil, errors.New("the fetchStackFrame func is mandatory in order to build a MachineLanguage instance")
	}

	if app.machine == nil {
		return nil, errors.New("the machine is mandatory in order to build a MachineLanguage instance")
	}

	if app.input == nil {
		app.input = map[string]computable.Value{}
	}

	root := app.lang.Root()
	paths := app.lang.Paths()

	patternMatches := map[string]definitions.PatternMatch{}
	patternMatchesList := app.lang.PatternMatches()
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

	lbls := map[string]label_instructions.Instructions{}
	languageLabels := app.lang.Application().Labels().All()
	for _, oneLabel := range languageLabels {
		name := oneLabel.Name()
		lbls[name] = oneLabel.Instructions()
	}

	lexerAdapterBuilder := app.lexerAdapterBuilder.Create().WithGrammarRetrieverCriteria(retrieverCriteria).WithEvents(app.events)
	return createMachineLanguage(
		app.variableBuilder,
		app.valueBuilder,
		app.computableValueBuilder,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		lexerAdapterBuilder,
		patternMatches,
		lbls,
		app.machine,
		app.fetchStackFrameFn,
	), nil
}
