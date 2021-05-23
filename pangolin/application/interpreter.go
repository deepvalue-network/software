package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
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
func (app *interpreter) Execute(excutable linkers.Executable, input map[string]computable.Value) (stackframes.StackFrame, error) {
	return app.in.Execute(excutable, input)
}

// Tests executes the tests of an executable
func (app *interpreter) Tests(excutable linkers.Executable) error {
	return app.in.Tests(excutable)
}
