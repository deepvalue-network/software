package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
)

type testable struct {
	testInsMachineBuilder     machines.TestInstructionBuilder
	machineStateFactory       machines.LanguageStateFactory
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	stackFrameBuilder         stackframes.Builder
	composerBuilder           composers.Builder
	executable                Executable
	linker                    linkers.Linker
	events                    []lexers.Event
}

func createTestable(
	testInsMachineBuilder machines.TestInstructionBuilder,
	machineStateFactory machines.LanguageStateFactory,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	stackFrameBuilder stackframes.Builder,
	composerBuilder composers.Builder,
	executable Executable,
	linker linkers.Linker,
	events []lexers.Event,
) Testable {
	out := testable{
		testInsMachineBuilder:     testInsMachineBuilder,
		machineStateFactory:       machineStateFactory,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		stackFrameBuilder:         stackFrameBuilder,
		composerBuilder:           composerBuilder,
		executable:                executable,
		linker:                    linker,
		events:                    events,
	}

	return &out
}

// Execute executes tests
func (app *testable) Execute(testable linkers.Testable) error {
	if testable.IsExecutable() {
		executable := testable.Executable()
		return app.Executable(executable)
	}

	langRef := testable.Language().Definition()
	return app.Language(langRef)
}

// Executable executes tests on an executable instance
func (app *testable) Executable(executable linkers.Executable) error {
	if executable.IsApplication() {
		appli := executable.Application()
		return app.Application(appli)
	}

	script := executable.Script()
	return app.Script(script)
}

// Application executes tests on an application instance
func (app *testable) Application(linkedApp linkers.Application) error {
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

// Script executes tests on a script application instance
func (app *testable) Script(script linkers.Script) error {
	// execute the language tests first:
	langDef := script.Language().Definition()
	err := app.Language(langDef)
	if err != nil {
		return err
	}

	// if there is no tests, return successfully:
	if !script.HasTests() {
		return nil
	}

	// execute the script tests:
	tests := script.Tests()
	fmt.Printf(delimiter)
	fmt.Printf("Executing %d script tests...\n", len(tests))
	fmt.Printf(delimiter)
	for index, oneTest := range tests {
		name := oneTest.Name()
		executable := oneTest.Executable()

		fmt.Printf(delimiter)
		fmt.Printf(printTestStr, name)
		input := app.stackFrameBuilder.Create().WithVariables(map[string]computable.Value{}).Now()
		_, err := app.executable.Execute(executable, input)
		if err != nil {
			str := fmt.Sprintf("error while executing script application... index: %d, error: %s", index, err.Error())
			return errors.New(str)
		}
	}

	fmt.Printf(delimiter)

	return nil
}

// Language executes tests on a language application instance
func (app *testable) Language(linkedLangDef linkers.LanguageDefinition) error {
	interpreterCallBackFn := func(composedApp linkers.Application, input stackframes.StackFrame) (stackframes.StackFrame, error) {
		return app.executable.Application(composedApp, input)
	}

	langApp := linkedLangDef.Application()
	tests := langApp.Tests().All()
	fmt.Printf(delimiter)
	fmt.Printf("Executing %d language tests...\n", len(tests))
	fmt.Printf(delimiter)

	for _, oneTest := range tests {
		languageState := app.machineStateFactory.Create()
		stackFrame := app.stackFrameBuilder.Create().Now()
		composer, err := app.composerBuilder.Create().WithStackFrame(stackFrame).WithLinker(app.linker).Now()
		if err != nil {
			return err
		}

		testInsMachine, err := app.machineLangTestInsBuilder.Create().
			WithComposer(composer).
			WithLanguage(linkedLangDef).
			WithStackFrame(stackFrame).
			WithState(languageState).
			WithEvents(app.events).
			WithInterpreterCallBackkFn(interpreterCallBackFn).
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
			if stackFrame.Current().IsStopped() {
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
