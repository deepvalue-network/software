package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type builder struct {
	stackFrameBuilder      StackFrameBuilder
	machineBuilder         MachineBuilder
	machineLanguageBuilder MachineLanguageBuilder
	valueBuilder           computable.Builder
	program                linkers.Program
}

func createBuilder(
	stackFrameBuilder StackFrameBuilder,
	machineBuilder MachineBuilder,
	machineLanguageBuilder MachineLanguageBuilder,
	valueBuilder computable.Builder,
) Builder {
	out := builder{
		stackFrameBuilder:      stackFrameBuilder,
		machineBuilder:         machineBuilder,
		machineLanguageBuilder: machineLanguageBuilder,
		valueBuilder:           valueBuilder,
		program:                nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.stackFrameBuilder,
		app.machineBuilder,
		app.machineLanguageBuilder,
		app.valueBuilder,
	)
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program linkers.Program) Builder {
	app.program = program
	return app
}

// Now builds a new Interpreter instance
func (app *builder) Now() (Interpreter, error) {
	if app.program == nil {
		return nil, errors.New("the linked Program is mandatory in order to build an Interpreter instance")
	}

	if app.program.IsScript() {
		linkedScript := app.program.Script()
		script := createScript(app.stackFrameBuilder, app.machineBuilder, app.machineLanguageBuilder, app.valueBuilder, linkedScript)
		return createInterpreterWithScript(script), nil
	}

	if app.program.IsApplication() {
		linkedApp := app.program.Application()
		app := createApplication(app.machineBuilder, app.stackFrameBuilder, linkedApp)
		return createInterpreterWithApplication(app), nil
	}

	if app.program.IsLanguage() {
		linkedLang := app.program.Language()
		if linkedLang.IsReference() {
			linkedLangDef := linkedLang.Reference().Definition()
			lang := createLanguage(app.stackFrameBuilder, app.machineBuilder, app.machineLanguageBuilder, app.valueBuilder, linkedLangDef)
			return createInterpreterWithLanguage(lang), nil
		}

		return nil, errors.New("->-> finish the language application in the builder")
	}

	return nil, errors.New("the Interpreter is invalid")
}
