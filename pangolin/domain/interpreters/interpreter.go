package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type interpreter struct {
	application Application
	script      Script
}

func createInterpreter(
	application Application,
	script Script,
) Interpreter {
	out := interpreter{
		application: application,
		script:      script,
	}

	return &out
}

// Execute executes the interpreter
func (app *interpreter) Execute(excutable linkers.Executable, input map[string]computable.Value) (stackframes.StackFrame, error) {
	if excutable.IsApplication() {
		linkedApp := excutable.Application()
		return app.application.Execute(linkedApp, input)
	}

	linkedScript := excutable.Script()
	linkedApp, err := app.script.Execute(linkedScript)
	if err != nil {
		return nil, err
	}

	return app.application.Execute(linkedApp, input)
}

// Tests executes the  tests
func (app *interpreter) Tests(excutable linkers.Executable) error {
	if excutable.IsApplication() {
		linkedApp := excutable.Application()
		return app.application.Tests(linkedApp)
	}

	linkedScript := excutable.Script()
	return app.script.Tests(linkedScript)
}
