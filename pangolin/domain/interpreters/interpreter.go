package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type interpreter struct {
	testable   Testable
	executable Executable
}

func createInterpreter(
	testable Testable,
	executable Executable,
) Interpreter {
	out := interpreter{
		testable:   testable,
		executable: executable,
	}

	return &out
}

// Execute executes an executable instance
func (app *interpreter) Execute(excutable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error) {
	return app.executable.Execute(excutable, input)
}

// Tests executes tests on a testable instance
func (app *interpreter) Tests(testable linkers.Testable) error {
	return app.testable.Execute(testable)
}
