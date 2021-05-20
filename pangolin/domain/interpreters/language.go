package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type language struct {
	composerBuilder           composers.Builder
	machineStateFactory       machines.LanguageStateFactory
	stackFrameBuilder         stackframes.Builder
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	machineLangInsBuilder     machines.LanguageInstructionBuilder
}

func createLanguage(
	composerBuilder composers.Builder,
	machineStateFactory machines.LanguageStateFactory,
	stackFrameBuilder stackframes.Builder,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	machineLangInsBuilder machines.LanguageInstructionBuilder,
) Language {
	out := language{
		composerBuilder:           composerBuilder,
		machineStateFactory:       machineStateFactory,
		stackFrameBuilder:         stackFrameBuilder,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		machineLangInsBuilder:     machineLangInsBuilder,
	}

	return &out
}

// Execute executes a language application
func (app *language) Execute(linkedLangDef linkers.LanguageDefinition, input map[string]computable.Value) (linkers.Application, error) {
	stackFrame := app.stackFrameBuilder.Create().WithVariables(input).Now()
	composer, err := app.composerBuilder.Create().WithStackFrame(stackFrame).Now()
	if err != nil {
		return nil, err
	}

	state := app.machineStateFactory.Create()
	machineLangInsApp, err := app.machineLangInsBuilder.Create().WithComposer(composer).WithLanguage(linkedLangDef).WithStackFrame(stackFrame).WithState(state).Now()
	if err != nil {
		return nil, err
	}

	insList := linkedLangDef.Application().Instructions().All()
	for _, oneIns := range insList {
		err := machineLangInsApp.Receive(oneIns)
		if err != nil {
			return nil, err
		}
	}

	return composer.Now()
}

//TestsAll executes all tests
func (app *language) TestsAll(linkedLangDef linkers.LanguageDefinition) error {
	names := []string{}
	tests := linkedLangDef.Application().Tests().All()
	for _, oneTest := range tests {
		name := oneTest.Name()
		names = append(names, name)
	}

	return app.TestByNames(linkedLangDef, names)
}

// TestByNames executes the tests of an application in the interpreter
func (app *language) TestByNames(linkedLangDef linkers.LanguageDefinition, names []string) error {
	fmt.Printf("\n++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("Executing %d language tests...\n", len(names))
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")

	langApp := linkedLangDef.Application()
	tests := langApp.Tests().All()
	for _, oneTest := range tests {
		languageState := app.machineStateFactory.Create()
		stackframe := app.stackFrameBuilder.Create().Now()
		testInsMachine, err := app.machineLangTestInsBuilder.Create().
			WithLanguage(linkedLangDef).
			WithStackFrame(stackframe).
			WithState(languageState).
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
