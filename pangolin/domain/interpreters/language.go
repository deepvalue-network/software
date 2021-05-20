package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type language struct {
	machineStateFactory       machines.LanguageStateFactory
	stackFrameBuilder         stackframes.Builder
	machineLanguageBuilder    machines.LanguageInstructionBuilder
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	languageDef               linkers.LanguageDefinition
}

func createLanguage(
	machineStateFactory machines.LanguageStateFactory,
	stackFrameBuilder stackframes.Builder,
	machineLanguageBuilder machines.LanguageInstructionBuilder,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	languageDef linkers.LanguageDefinition,
) Language {
	out := language{
		machineStateFactory:       machineStateFactory,
		stackFrameBuilder:         stackFrameBuilder,
		machineLanguageBuilder:    machineLanguageBuilder,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		languageDef:               languageDef,
	}

	return &out
}

//TestsAll executes all tests
func (app *language) TestsAll() error {
	names := []string{}
	tests := app.languageDef.Application().Tests().All()
	for _, oneTest := range tests {
		name := oneTest.Name()
		names = append(names, name)
	}

	return app.TestByNames(names)
}

// TestByNames executes the tests of an application in the interpreter
func (app *language) TestByNames(names []string) error {
	fmt.Printf("\n++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("Executing %d language tests...\n", len(names))
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")

	baseDir := app.languageDef.Paths().BaseDir()
	langApp := app.languageDef.Application()
	tests := langApp.Tests().All()
	labels := langApp.Labels()
	for _, oneTest := range tests {
		languageState := app.machineStateFactory.Create()
		stackframe := app.stackFrameBuilder.Create().Now()
		testInsMachine, err := app.machineLangTestInsBuilder.Create().
			WithStackFrame(stackframe).
			WithLabels(labels).
			WithState(languageState).
			WithBaseDir(baseDir).
			Now()

		if err != nil {
			return err
		}

		name := oneTest.Name()
		fmt.Printf("\n-----------------------------------\n")
		fmt.Printf("Test: %s\n", name)
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

		fmt.Printf("-----------------------------------\n")
	}

	return nil
}
