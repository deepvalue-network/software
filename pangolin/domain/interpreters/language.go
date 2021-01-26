package interpreters

import (
	"fmt"

	"github.com/steve-care-software/products/pangolin/domain/linkers"
)

type language struct {
	machineBuilder MachineBuilder
	language       linkers.Language
}

func createLanguage(
	machineBuilder MachineBuilder,
	linkedLanguage linkers.Language,
) Language {
	out := language{
		machineBuilder: machineBuilder,
		language:       linkedLanguage,
	}

	return &out
}

//TestsAll executes all tests
func (app *language) TestsAll() error {
	names := []string{}
	tests := app.language.Application().Tests().All()
	for _, oneTest := range tests {
		name := oneTest.Name()
		names = append(names, name)
	}

	return app.Tests(names)
}

// Tests executes the tests of an application in the interpreter
func (app *language) Tests(names []string) error {
	machine, err := app.machineBuilder.Create().WithLanguage(app.language).Now()
	if err != nil {
		return err
	}

	fmt.Printf("\n++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("Executing %d tests...\n", len(names))
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")

	tests := app.language.Application().Tests().All()
	for _, oneTest := range tests {
		name := oneTest.Name()
		fmt.Printf("\n-----------------------------------\n")
		fmt.Printf("Test: %s\n", name)
		testInstructions := oneTest.Instructions().All()
		for _, oneTestInstruction := range testInstructions {
			if oneTestInstruction.IsStart() {
				fmt.Println("Begins.")
				continue
			}

			if oneTestInstruction.IsStop() {
				fmt.Println("Ends.")
				continue
			}

			if oneTestInstruction.IsInstruction() {
				ins := oneTestInstruction.Instruction()
				err := machine.Receive(ins)
				if err != nil {
					return err
				}
			}
		}

		fmt.Printf("-----------------------------------\n")
	}

	return nil
}
