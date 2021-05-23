package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type application struct {
	insMachineBuilder     machines.InstructionBuilder
	testInsMachineBuilder machines.TestInstructionBuilder
	stackFrameBuilder     stackframes.Builder
}

func createApplication(
	insMachineBuilder machines.InstructionBuilder,
	testInsMachineBuilder machines.TestInstructionBuilder,
	stackFrameBuilder stackframes.Builder,
) Application {
	out := application{
		insMachineBuilder:     insMachineBuilder,
		testInsMachineBuilder: testInsMachineBuilder,
		stackFrameBuilder:     stackFrameBuilder,
	}

	return &out
}

// Execute executes an application in the interpreter
func (app *application) Execute(linkedApp linkers.Application, input map[string]computable.Value) (stackframes.StackFrame, error) {
	labels := linkedApp.Labels()
	stackFrame := app.stackFrameBuilder.Create().WithVariables(input).Now()
	machine, err := app.insMachineBuilder.Create().WithLabels(labels).WithStackFrame(stackFrame).Now()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	err = app.execute(machine, linkedApp)
	if err != nil {
		return nil, err
	}

	return stackFrame, nil
}

// Tests execute tests
func (app *application) Tests(linkedApp linkers.Application) error {
	tests := linkedApp.Tests().All()
	fmt.Printf(delimiter)
	fmt.Printf("Executing %d application tests...\n", len(tests))
	fmt.Printf(delimiter)

	baseDir := "./"
	labels := linkedApp.Labels()
	for _, oneTest := range tests {
		stackframe := app.stackFrameBuilder.Create().Now()
		testInsMachine, err := app.testInsMachineBuilder.Create().
			WithStackFrame(stackframe).
			WithLabels(labels).
			WithBaseDir(baseDir).
			Now()

		if err != nil {
			return err
		}

		name := oneTest.Name()
		fmt.Printf(delimiter)
		fmt.Printf(printTestStr, name)
		testInstructions := oneTest.Instructions().All()
		for index, oneTestInstruction := range testInstructions {
			// if the machine is stopped, stop:
			if stackframe.Current().IsStopped() {
				return nil
			}

			stops, err := testInsMachine.Receive(oneTestInstruction)
			if err != nil {
				str := fmt.Sprintf("index: %d, error: %s", index, err.Error())
				return errors.New(str)
			}

			if stops {
				break
			}
		}

		fmt.Printf(delimiter)
	}

	return nil
}

func (app *application) execute(machine machines.Instruction, application linkers.Application) error {
	ins := application.Instructions().All()
	for _, oneIns := range ins {
		err := machine.Receive(oneIns)
		if err != nil {
			return err
		}
	}

	return nil
}
