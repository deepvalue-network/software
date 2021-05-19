package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	language_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
)

type languageTestInstructionBuilder struct {
	insLangCommonApp  LanguageInstructionCommonBuilder
	testInsAppBuilder TestInstructionBuilder
	stackFrame        stackframes.StackFrame
	labels            labels.Labels
	langLabels        language_labels.Labels
	state             LanguageState
	baseDir           string
}

func createLanguageTestInstructionBuilder(
	insLangCommonApp LanguageInstructionCommonBuilder,
	testInsAppBuilder TestInstructionBuilder,
) LanguageTestInstructionBuilder {
	out := languageTestInstructionBuilder{
		insLangCommonApp:  insLangCommonApp,
		testInsAppBuilder: testInsAppBuilder,
		stackFrame:        nil,
		labels:            nil,
		langLabels:        nil,
		state:             nil,
		baseDir:           "",
	}

	return &out
}

// Create initializes the builder
func (app *languageTestInstructionBuilder) Create() LanguageTestInstructionBuilder {
	return createLanguageTestInstructionBuilder(app.insLangCommonApp, app.testInsAppBuilder)
}

// WithStackFrame adds a stackFrame to the builder
func (app *languageTestInstructionBuilder) WithStackFrame(stackFrame stackframes.StackFrame) LanguageTestInstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithLabels add labels to the builder
func (app *languageTestInstructionBuilder) WithLabels(labels labels.Labels) LanguageTestInstructionBuilder {
	app.labels = labels
	return app
}

// WithLanguageLabels add language labels to the builder
func (app *languageTestInstructionBuilder) WithLanguageLabels(langLabels language_labels.Labels) LanguageTestInstructionBuilder {
	app.langLabels = langLabels
	return app
}

// WithState adds a state to the builder
func (app *languageTestInstructionBuilder) WithState(state LanguageState) LanguageTestInstructionBuilder {
	app.state = state
	return app
}

// WithBaseDir adds a base directory to the builder
func (app *languageTestInstructionBuilder) WithBaseDir(baseDir string) LanguageTestInstructionBuilder {
	app.baseDir = baseDir
	return app
}

// Now builds a new LanguageTestInstruction instance
func (app *languageTestInstructionBuilder) Now() (LanguageTestInstruction, error) {
	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.labels == nil {
		return nil, errors.New("the Labels are mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.langLabels == nil {
		return nil, errors.New("the language Labels are mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.state == nil {
		return nil, errors.New("the LanguageState is mandatory in order to build a LanguageTestInstruction instance")
	}

	if app.baseDir == "" {
		return nil, errors.New("the base directory is mandatory in order to build a LanguageTestInstruction instance")
	}

	langCommonInsApp, err := app.insLangCommonApp.Create().WithState(app.state).WithLabels(app.langLabels).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	testInsApp, err := app.testInsAppBuilder.Create().WithLabels(app.labels).WithStackFrame(app.stackFrame).WithBaseDir(app.baseDir).Now()
	if err != nil {
		return nil, err
	}

	return createLanguageTestInstruction(langCommonInsApp, testInsApp), nil
}
