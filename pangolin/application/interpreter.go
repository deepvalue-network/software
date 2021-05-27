package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type interpreter struct {
	in interpreters.Interpreter
}

func createInterpreter(
	in interpreters.Interpreter,
) Interpreter {
	out := interpreter{
		in: in,
	}

	return &out
}

// Execute executes an executable
func (app *interpreter) Execute(excutable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error) {
	return app.in.Execute(excutable, input)
}

// Tests executes the tests of an executable
func (app *interpreter) Tests(testable linkers.Testable) error {
	return app.in.Tests(testable)
}
