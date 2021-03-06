package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/labels/label"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/labels/label/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type machineBuilder struct {
	computableValueBuilder          computable.Builder
	lexerParserApplication          lexer_parser.Application
	lexerParserBuilder              lexer_parser.Builder
	lexerAdapterBuilder             lexers.AdapterBuilder
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	stackFrameBuilder               StackFrameBuilder
	events                          []lexers.Event
	globalConstants                 map[string]computable.Value
	globalVariables                 map[string]computable.Value
	lbls                            map[string]label_instructions.Instructions
	lang                            linkers.Language
	app                             linkers.Application
	input                           map[string]computable.Value
}

func createMachineBuilder(
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerAdapterBuilder lexers.AdapterBuilder,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	stackFrameBuilder StackFrameBuilder,
	events []lexers.Event,
) MachineBuilder {
	out := machineBuilder{
		computableValueBuilder:          computableValueBuilder,
		lexerParserApplication:          lexerParserApplication,
		lexerParserBuilder:              lexerParserBuilder,
		lexerAdapterBuilder:             lexerAdapterBuilder,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		stackFrameBuilder:               stackFrameBuilder,
		events:                          events,
		globalConstants:                 map[string]computable.Value{},
		globalVariables:                 map[string]computable.Value{},
		lbls:                            map[string]label_instructions.Instructions{},
		lang:                            nil,
		app:                             nil,
		input:                           nil,
	}

	return &out
}

// Create initializes the builder
func (app *machineBuilder) Create() MachineBuilder {
	return createMachineBuilder(
		app.computableValueBuilder,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerAdapterBuilder,
		app.grammarRetrieverCriteriaBuilder,
		app.stackFrameBuilder,
		app.events,
	)
}

// WithLanguage adds a language to the builder
func (app *machineBuilder) WithLanguage(lang linkers.Language) MachineBuilder {
	app.lang = lang
	return app
}

// WithApplication adds an application to the builder
func (app *machineBuilder) WithApplication(appli linkers.Application) MachineBuilder {
	app.app = appli
	return app
}

// WithInput adds input values to the builder
func (app *machineBuilder) WithInput(input map[string]computable.Value) MachineBuilder {
	app.input = input
	return app
}

// Now builds a new Machine instance
func (app *machineBuilder) Now() (Machine, error) {
	if app.lang != nil {
		app.app = app.lang.Application()
	}

	if app.app == nil {
		return nil, errors.New("the application is mandatory in order to build a Machine instance")
	}

	if app.input == nil {
		app.input = map[string]computable.Value{}
	}

	variables := app.app.Variables()
	err := app.variables(variables)
	if err != nil {
		return nil, err
	}

	labels := app.app.Labels()
	err = app.labels(labels)
	if err != nil {
		return nil, err
	}

	stackFrame := app.stackFrameBuilder.Create().
		WithConstants(app.globalConstants).
		WithVariables(app.globalVariables).
		Now()

	if app.lang != nil {
		root := app.lang.Root()
		paths := app.lang.Paths()

		patternMatches := map[string]middle.PatternMatch{}
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

		lexerAdapterBuilder := app.lexerAdapterBuilder.Create().WithGrammarRetrieverCriteria(retrieverCriteria).WithEvents(app.events)
		return createMachineWithPatternMatches(app.computableValueBuilder, app.lexerParserApplication, app.lexerParserBuilder, stackFrame, app.lbls, patternMatches, lexerAdapterBuilder), nil
	}

	return createMachine(app.computableValueBuilder, app.lexerParserApplication, app.lexerParserBuilder, stackFrame, app.lbls), nil
}

func (app *machineBuilder) variables(variables variables.Variables) error {
	lst := variables.All()
	for _, oneVariable := range lst {
		err := app.variable(oneVariable)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *machineBuilder) variable(variable variable.Variable) error {
	name := variable.Name()
	if !variable.IsGlobal() {
		str := fmt.Sprintf("the application's variable (%s) was expected to be global", name)
		return errors.New(str)
	}

	isImmutable := false
	if variable.IsImmutable() {
		isImmutable = true
	}

	// if the variable has an incoming:
	if variable.IsIncoming() {
		// if there is an input with that name:
		if inputVal, ok := app.input[name]; ok {
			if isImmutable {
				app.globalConstants[name] = inputVal
				return nil
			}

			app.globalVariables[name] = inputVal
			return nil
		}
	}

	val := variable.Value()
	compValue, err := app.fetchComputableValue(val, isImmutable)
	if err != nil {
		return err
	}

	if isImmutable {
		app.globalConstants[name] = compValue
		return nil
	}

	app.globalVariables[name] = compValue
	return nil
}

func (app *machineBuilder) fetchComputableValue(val var_value.Value, isImmutable bool) (computable.Value, error) {
	if val.IsComputable() {
		return val.Computable(), nil
	}

	if val.IsGlobalVariable() {
		global := val.GlobalVariable()
		if isImmutable {
			if val, ok := app.globalConstants[global]; ok {
				return val, nil
			}

			str := fmt.Sprintf("the global constant (%s) is not declared", global)
			return nil, errors.New(str)
		}

		if val, ok := app.globalVariables[global]; ok {
			return val, nil
		}

		str := fmt.Sprintf("the global constant (%s) is not declared", global)
		return nil, errors.New(str)
	}

	return nil, errors.New("the given value is invalid")
}

func (app *machineBuilder) labels(labels labels.Labels) error {
	lst := labels.All()
	for _, oneLabel := range lst {
		err := app.label(oneLabel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *machineBuilder) label(lbl label.Label) error {
	name := lbl.Name()
	app.lbls[name] = lbl.Instructions()
	return nil
}
