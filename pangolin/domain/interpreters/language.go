package interpreters

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type language struct {
	stackFrameBuilder      StackFrameBuilder
	machineBuilder         MachineBuilder
	machineLanguageBuilder MachineLanguageBuilder
	valueBuilder           computable.Builder
	language               linkers.LanguageDefinition
}

func createLanguage(
	stackFrameBuilder StackFrameBuilder,
	machineBuilder MachineBuilder,
	machineLanguageBuilder MachineLanguageBuilder,
	valueBuilder computable.Builder,
	linkedLanguage linkers.LanguageDefinition,
) Language {
	out := language{
		stackFrameBuilder:      stackFrameBuilder,
		machineBuilder:         machineBuilder,
		machineLanguageBuilder: machineLanguageBuilder,
		valueBuilder:           valueBuilder,
		language:               linkedLanguage,
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
	fmt.Printf("\n++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("Executing %d tests...\n", len(names))
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")

	baseDir := app.language.Paths().BaseDir()
	langApp := app.language.Application()
	tests := langApp.Tests().All()
	labels := langApp.Labels()
	for _, oneTest := range tests {
		stackframe := app.stackFrameBuilder.Create().Now()
		machine, err := createMachineFromLanguageLabels(app.machineBuilder, stackframe, labels)
		if err != nil {
			return err
		}

		/*machineLanguage, err := app.machineLanguageBuilder.Create().WithLanguage(app.language).WithMachine(machine).Now()
		if err != nil {
			return err
		}*/

		name := oneTest.Name()
		fmt.Printf("\n-----------------------------------\n")
		fmt.Printf("Test: %s\n", name)
		testInstructions := oneTest.Instructions().All()
		for index, oneTestInstruction := range testInstructions {
			// if the machine is stopped, stop:
			if stackframe.Current().IsStopped() {
				break
			}

			if !oneTestInstruction.IsTest() {
				continue
			}

			testInstruction := oneTestInstruction.Test()
			if testInstruction.IsAssert() {
				assert := testInstruction.Assert()
				assertIndex := assert.Index()
				if assert.HasCondition() {
					condition := assert.Condition()
					condVal, err := stackframe.Current().Fetch(condition)
					if err != nil {
						return err
					}

					if !condVal.IsBool() {
						str := fmt.Sprintf("the assert's condition was expected to contain a bool, index: %d", index)
						return errors.New(str)
					}

					val := condVal.Bool()
					if *val {
						fmt.Printf("-> Assert, index: %d\n", assertIndex)
						break
					}

					continue
				}

				fmt.Printf("-> Assert, index: %d\n", assertIndex)
				break
			}

			if testInstruction.IsInstruction() {
				ins := testInstruction.Instruction()
				err := machine.Receive(ins)
				if err != nil {
					return err
				}

				continue
			}

			if testInstruction.IsReadFile() {
				readFile := testInstruction.ReadFile()
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
				err = stackframe.Current().UpdateValue(variable, computable)
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
