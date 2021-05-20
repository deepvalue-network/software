package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type languageInstructionBuilder struct {
	variableBuilder        var_variable.Builder
	valueBuilder           var_value.Builder
	computableValueBuilder computable.Builder
	langCommonInsBuilder   LanguageInstructionCommonBuilder
	insAppBuilder          InstructionBuilder
	composerApp            composers.Composer
	langDef                linkers.LanguageDefinition
	stackFrame             stackframes.StackFrame
	state                  LanguageState
}

func createLanguageInstructionBuilder(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	langCommonInsBuilder LanguageInstructionCommonBuilder,
	insAppBuilder InstructionBuilder,
) LanguageInstructionBuilder {
	out := languageInstructionBuilder{
		variableBuilder:        variableBuilder,
		valueBuilder:           valueBuilder,
		computableValueBuilder: computableValueBuilder,
		langCommonInsBuilder:   langCommonInsBuilder,
		insAppBuilder:          insAppBuilder,
		composerApp:            nil,
		langDef:                nil,
		stackFrame:             nil,
		state:                  nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageInstructionBuilder) Create() LanguageInstructionBuilder {
	return createLanguageInstructionBuilder(
		app.variableBuilder,
		app.valueBuilder,
		app.computableValueBuilder,
		app.langCommonInsBuilder,
		app.insAppBuilder,
	)
}

// WithComposer adds a composer to the builder
func (app *languageInstructionBuilder) WithComposer(composerApp composers.Composer) LanguageInstructionBuilder {
	app.composerApp = composerApp
	return app
}

// WithLanguage adds a language definition to the builder
func (app *languageInstructionBuilder) WithLanguage(langDef linkers.LanguageDefinition) LanguageInstructionBuilder {
	app.langDef = langDef
	return app
}

// WithStackFrame adds a stackframe to the builder
func (app *languageInstructionBuilder) WithStackFrame(stackFrame stackframes.StackFrame) LanguageInstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithState adds a state to the builder
func (app *languageInstructionBuilder) WithState(state LanguageState) LanguageInstructionBuilder {
	app.state = state
	return app
}

// Now builds a new LanguageInstruction instance
func (app *languageInstructionBuilder) Now() (LanguageInstruction, error) {
	if app.composerApp == nil {
		return nil, errors.New("the composer is mandatory in order to build a LanguageInstruction instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a LanguageInstruction instance")
	}

	if app.langDef == nil {
		return nil, errors.New("the language definition is mandatory in order to build a LanguageInstruction instance")
	}

	if app.state == nil {
		return nil, errors.New("the state is mandatory in order to build a LanguageInstruction instance")
	}

	labels := app.langDef.Application().Labels()
	fn, err := fromLanguageLabelsToCallLabelByNameFunc(app.stackFrame, labels)
	if err != nil {
		return nil, err
	}

	insApp, err := app.insAppBuilder.Create().WithCallLabelFn(fn).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	langCommonIns, err := app.langCommonInsBuilder.Create().WithCallLabelFn(fn).WithLanguage(app.langDef).Now()
	if err != nil {
		return nil, err
	}

	return createLanguageInstruction(
		app.variableBuilder,
		app.valueBuilder,
		app.computableValueBuilder,
		langCommonIns,
		insApp,
		app.stackFrame,
		app.state,
		app.composerApp,
	), nil
}
