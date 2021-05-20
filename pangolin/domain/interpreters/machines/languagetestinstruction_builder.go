package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type languageTestInstructionBuilder struct {
	insLangCommonApp  LanguageInstructionCommonBuilder
	testInsAppBuilder TestInstructionBuilder
	langDef           linkers.LanguageDefinition
	stackFrame        stackframes.StackFrame
	state             LanguageState
}

func createLanguageTestInstructionBuilder(
	insLangCommonApp LanguageInstructionCommonBuilder,
	testInsAppBuilder TestInstructionBuilder,
) LanguageTestInstructionBuilder {
	out := languageTestInstructionBuilder{
		insLangCommonApp:  insLangCommonApp,
		testInsAppBuilder: testInsAppBuilder,
		langDef:           nil,
		stackFrame:        nil,
		state:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageTestInstructionBuilder) Create() LanguageTestInstructionBuilder {
	return createLanguageTestInstructionBuilder(app.insLangCommonApp, app.testInsAppBuilder)
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

// Now builds a new LanguageTestInstruction instance
func (app *languageTestInstructionBuilder) Now() (LanguageTestInstruction, error) {
	if app.langDef == nil {
		return nil, errors.New("the language definition is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.state == nil {
		return nil, errors.New("the LanguageState is mandatory in order to build a LanguageTestInstruction instance")
	}

	labels := app.langDef.Application().Labels()
	labelCallFn, err := fromLanguageLabelsToCallLabelByNameFunc(app.stackFrame, labels)
	if err != nil {
		return nil, err
	}

	langCommonInsApp, err := app.insLangCommonApp.Create().WithLanguage(app.langDef).WithState(app.state).WithCallLabelFn(labelCallFn).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	baseDir := app.langDef.Paths().BaseDir()
	testInsApp, err := app.testInsAppBuilder.Create().WithCallLabelFn(labelCallFn).WithStackFrame(app.stackFrame).WithBaseDir(baseDir).Now()
	if err != nil {
		return nil, err
	}

	return createLanguageTestInstruction(langCommonInsApp, testInsApp), nil
}
