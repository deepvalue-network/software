package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type application struct {
	machineBuilder MachineBuilder
	application    linkers.Application
}

func createApplication(
	machineBuilder MachineBuilder,
	linkedApplication linkers.Application,
) Application {
	out := application{
		machineBuilder: machineBuilder,
		application:    linkedApplication,
	}

	return &out
}

// Execute executes an application in the interpreter
func (app *application) Execute(input map[string]computable.Value) (StackFrame, error) {
	machine, err := app.machineBuilder.Create().WithInput(input).WithApplication(app.application).Now()
	if err != nil {
		return nil, err
	}

	return execute(machine, app.application)
}
