package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type language struct {
	composerBuilder           composers.Builder
	machineStateFactory       machines.LanguageStateFactory
	stackFrameBuilder         stackframes.Builder
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	machineLangInsBuilder     machines.LanguageInstructionBuilder
	linker                    linkers.Linker
	events                    []lexers.Event
}

func createLanguage(
	composerBuilder composers.Builder,
	machineStateFactory machines.LanguageStateFactory,
	stackFrameBuilder stackframes.Builder,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	machineLangInsBuilder machines.LanguageInstructionBuilder,
	linker linkers.Linker,
	events []lexers.Event,
) Language {
	out := language{
		composerBuilder:           composerBuilder,
		machineStateFactory:       machineStateFactory,
		stackFrameBuilder:         stackFrameBuilder,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		machineLangInsBuilder:     machineLangInsBuilder,
		linker:                    linker,
		events:                    events,
	}

	return &out
}

// Execute executes a language application
func (app *language) Execute(linkedLangDef linkers.LanguageDefinition, input map[string]computable.Value) (linkers.Application, error) {
	stackFrame := app.stackFrameBuilder.Create().WithVariables(input).Now()
	composer, err := app.composerBuilder.Create().WithStackFrame(stackFrame).WithLinker(app.linker).Now()
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

//Tests executes tests
func (app *language) Tests(linkedLangDef linkers.LanguageDefinition) error {
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
