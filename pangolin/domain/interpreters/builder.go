package interpreters

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/linkers"
	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value/computable"
)

type builder struct {
	machineBuilder MachineBuilder
	valueBuilder   computable.Builder
	program        linkers.Program
}

func createBuilder(
	machineBuilder MachineBuilder,
	valueBuilder computable.Builder,
) Builder {
	out := builder{
		machineBuilder: machineBuilder,
		valueBuilder:   valueBuilder,
		program:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.machineBuilder, app.valueBuilder)
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
		script := createScript(app.machineBuilder, app.valueBuilder, linkedScript)
		return createInterpreterWithScript(script), nil
	}

	if app.program.IsApplication() {
		linkedApp := app.program.Application()
		app := createApplication(app.machineBuilder, linkedApp)
		return createInterpreterWithApplication(app), nil
	}

	if app.program.IsLanguage() {
		linkedLang := app.program.Language().Language()
		lang := createLanguage(app.machineBuilder, app.valueBuilder, linkedLang)
		return createInterpreterWithLanguage(lang), nil
	}

	return nil, errors.New("the Interpreter is invalid")
}
