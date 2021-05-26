package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type languageTestInstructionBuilder struct {
	frameBuilder          stackframes.FrameBuilder
	insLangCommonApp      LanguageInstructionCommonBuilder
	testInsAppBuilder     TestInstructionBuilder
	langInsBuilder        LanguageInstructionBuilder
	interpreterCallBackFn InterpretCallBackFn
	composerApp           composers.Composer
	langDef               linkers.LanguageDefinition
	stackFrame            stackframes.StackFrame
	state                 LanguageState
	events                []lexers.Event
}

func createLanguageTestInstructionBuilder(
	frameBuilder stackframes.FrameBuilder,
	insLangCommonApp LanguageInstructionCommonBuilder,
	testInsAppBuilder TestInstructionBuilder,
	langInsBuilder LanguageInstructionBuilder,
) LanguageTestInstructionBuilder {
	out := languageTestInstructionBuilder{
		frameBuilder:          frameBuilder,
		insLangCommonApp:      insLangCommonApp,
		testInsAppBuilder:     testInsAppBuilder,
		langInsBuilder:        langInsBuilder,
		interpreterCallBackFn: nil,
		composerApp:           nil,
		langDef:               nil,
		stackFrame:            nil,
		state:                 nil,
		events:                nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageTestInstructionBuilder) Create() LanguageTestInstructionBuilder {
	return createLanguageTestInstructionBuilder(app.frameBuilder, app.insLangCommonApp, app.testInsAppBuilder, app.langInsBuilder)
}

// WithInterpreterCallBackkFn adds an interpreter callbackfn to the builder
func (app *languageTestInstructionBuilder) WithInterpreterCallBackkFn(interpreterCallBackFn InterpretCallBackFn) LanguageTestInstructionBuilder {
	app.interpreterCallBackFn = interpreterCallBackFn
	return app
}

// WithComposer adds a composer to the builder
func (app *languageTestInstructionBuilder) WithComposer(composerApp composers.Composer) LanguageTestInstructionBuilder {
	app.composerApp = composerApp
	return app
}

// WithLanguage adds a language definition to the builder
func (app *languageTestInstructionBuilder) WithLanguage(langDef linkers.LanguageDefinition) LanguageTestInstructionBuilder {
	app.langDef = langDef
	return app
}

// WithStackFrame adds a stackFrame to the builder
func (app *languageTestInstructionBuilder) WithStackFrame(stackFrame stackframes.StackFrame) LanguageTestInstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithState adds a state to the builder
func (app *languageTestInstructionBuilder) WithState(state LanguageState) LanguageTestInstructionBuilder {
	app.state = state
	return app
}

// WithEvents add events to the builder
func (app *languageTestInstructionBuilder) WithEvents(events []lexers.Event) LanguageTestInstructionBuilder {
	app.events = events
	return app
}

// Now builds a new LanguageTestInstruction instance
func (app *languageTestInstructionBuilder) Now() (LanguageTestInstruction, error) {
	if app.interpreterCallBackFn == nil {
		return nil, errors.New("the interpreter callBack func is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.composerApp == nil {
		return nil, errors.New("the composer is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.langDef == nil {
		return nil, errors.New("the language definition is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.state == nil {
		return nil, errors.New("the LanguageState is mandatory in order to build a LanguageTestInstruction instance")
	}

	langInsApp, err := app.langInsBuilder.Create().
		WithComposer(app.composerApp).
		WithLanguage(app.langDef).
		WithStackFrame(app.stackFrame).
		WithState(app.state).
		WithEvents(app.events).
		Now()

	if err != nil {
		return nil, err
	}

	labels := app.langDef.Application().Labels()
	labelCallFn, err := fromLanguageLabelsToCallLabelByNameFunc(langInsApp, app.stackFrame, labels)
	if err != nil {
		return nil, err
	}

	langCommonInsApp, err := app.insLangCommonApp.Create().
		WithLanguage(app.langDef).
		WithState(app.state).
		WithCallLabelFn(labelCallFn).
		WithStackFrame(app.stackFrame).
		WithEvents(app.events).
		Now()

	if err != nil {
		return nil, err
	}

	baseDir := app.langDef.Paths().BaseDir()
	testInsApp, err := app.testInsAppBuilder.Create().
		WithCallLabelFn(labelCallFn).
		WithStackFrame(app.stackFrame).
		WithBaseDir(baseDir).
		Now()

	if err != nil {
		return nil, err
	}

	return createLanguageTestInstruction(
		app.frameBuilder,
		langCommonInsApp,
		testInsApp,
		app.composerApp,
		app.stackFrame,
		app.interpreterCallBackFn,
	), nil
}
