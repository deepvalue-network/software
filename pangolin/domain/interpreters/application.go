package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type application struct {
	machineBuilder    MachineBuilder
	stackFrameBuilder StackFrameBuilder
	application       linkers.Application
}

func createApplication(
	machineBuilder MachineBuilder,
	stackFrameBuilder StackFrameBuilder,
	linkedApplication linkers.Application,
) Application {
	out := application{
		machineBuilder:    machineBuilder,
		stackFrameBuilder: stackFrameBuilder,
		application:       linkedApplication,
	}

	return &out
}

// Execute executes an application in the interpreter
func (app *application) Execute(input map[string]computable.Value) (StackFrame, error) {
	labels := app.application.Labels()
	stackFrame := app.stackFrameBuilder.Create().WithVariables(input).Now()
	machine, err := createMachineFromLabels(
		app.machineBuilder,
		stackFrame,
		labels,
	)

	if err != nil {
		return nil, err
	}

	err = app.execute(machine, app.application)
	if err != nil {
		return nil, err
	}

	return stackFrame, nil
}

func (app *application) execute(machine Machine, application linkers.Application) error {
	ins := application.Instructions().All()
	for _, oneIns := range ins {
		err := machine.Receive(oneIns)
		if err != nil {
			return err
		}
	}

	return nil
}
