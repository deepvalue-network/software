package interpreters

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type language struct {
	machineBuilder MachineBuilder
	valueBuilder   computable.Builder
	language       linkers.Language
}

func createLanguage(
	machineBuilder MachineBuilder,
	valueBuilder computable.Builder,
	linkedLanguage linkers.Language,
) Language {
	out := language{
		machineBuilder: machineBuilder,
		valueBuilder:   valueBuilder,
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

	baseDir := app.language.Paths().BaseDir()
	tests := app.language.Application().Tests().All()
	for _, oneTest := range tests {
		name := oneTest.Name()
		fmt.Printf("\n-----------------------------------\n")
		fmt.Printf("Test: %s\n", name)
		testInstructions := oneTest.Instructions().All()
		for _, oneTestInstruction := range testInstructions {
			// if the machine is stopped, stop:
			if machine.StackFrame().Current().IsStopped() {
				break
			}

			if oneTestInstruction.IsAssert() {
				fmt.Printf("-> Assert !!\n")
				break
			}

			if oneTestInstruction.IsInstruction() {
				ins := oneTestInstruction.Instruction()
				err := machine.Receive(ins)
				if err != nil {
					return err
				}

				continue
			}

			if oneTestInstruction.IsReadFile() {
				readFile := oneTestInstruction.ReadFile()
				relativePath := readFile.Path()
				joinedPath := filepath.Join(baseDir, relativePath)
				absPath, err := filepath.Abs(joinedPath)
				if err != nil {
					str := fmt.Sprintf("there was an error while reading the relative path (%s): %s", relativePath, err.Error())
					return errors.New(str)
				}

				contents, err := ioutil.ReadFile(absPath)
				if err != nil {
					return err
				}

				contentsStr := string(contents)
				computable, err := app.valueBuilder.Create().WithString(contentsStr).Now()
				if err != nil {
					return err
				}

				variable := readFile.Variable()
				err = machine.StackFrame().Current().UpdateValue(variable, computable)
				if err != nil {
					return err
				}

				continue
			}
		}

		fmt.Printf("-----------------------------------\n")
	}

	return nil
}
