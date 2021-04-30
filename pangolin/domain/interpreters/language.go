package interpreters

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	language_test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
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
		machine, _, fetchStackFrameFunc, err := createMachineFromLanguageLabels(app.machineBuilder, stackframe, labels)
		if err != nil {
			return err
		}

		machineLanguage, err := app.machineLanguageBuilder.Create().WithLanguage(app.language).WithMachine(machine).WithFetchStackFunc(fetchStackFrameFunc).Now()
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

			stops, err := app.languageTestInstruction(oneTestInstruction, baseDir, index, stackframe, machine, machineLanguage)
			if err != nil {
				return err
			}

			if stops {
				break
			}
		}

		fmt.Printf("-----------------------------------\n")
	}

	return nil
}

func (app *language) languageTestInstruction(
	testIns language_test_instruction.Instruction,
	baseDir string,
	index int,
	stackframe StackFrame,
	machine Machine,
	machineLanguage MachineLanguage,
) (bool, error) {
	if testIns.IsLanguage() {
		langIns := testIns.Language()
		return app.languageInstruction(
			langIns,
			baseDir,
			index,
			stackframe,
			machine,
			machineLanguage,
		)
	}

	testInstruction := testIns.Test()
	return app.testInstruction(
		testInstruction,
		baseDir,
		index,
		stackframe,
		machine,
	)
}

func (app *language) testInstruction(
	testIns test_instruction.Instruction,
	baseDir string,
	index int,
	stackframe StackFrame,
	machine Machine,
) (bool, error) {
	if testIns.IsAssert() {
		assert := testIns.Assert()
		assertIndex := assert.Index()
		if assert.HasCondition() {
			condition := assert.Condition()
			condVal, err := stackframe.Current().Fetch(condition)
			if err != nil {
				return false, err
			}

			if !condVal.IsBool() {
				str := fmt.Sprintf("the assert's condition was expected to contain a bool, index: %d", index)
				return false, errors.New(str)
			}

			val := condVal.Bool()
			if *val {
				fmt.Printf("-> Assert, index: %d\n", assertIndex)
				return true, nil
			}

			return false, nil
		}

		fmt.Printf("-> Assert, index: %d\n", assertIndex)
		return true, nil
	}

	if testIns.IsInstruction() {
		ins := testIns.Instruction()
		return app.instruction(ins, machine)
	}

	if testIns.IsReadFile() {
		readFile := testIns.ReadFile()
		relativePath := readFile.Path()
		joinedPath := filepath.Join(baseDir, relativePath)
		absPath, err := filepath.Abs(joinedPath)
		if err != nil {
			str := fmt.Sprintf("there was an error while reading the relative path (%s): %s", relativePath, err.Error())
			return false, errors.New(str)
		}

		contents, err := ioutil.ReadFile(absPath)
		if err != nil {
			return false, err
		}

		contentsStr := string(contents)
		computable, err := app.valueBuilder.Create().WithString(contentsStr).Now()
		if err != nil {
			return false, err
		}

		variable := readFile.Variable()
		err = stackframe.Current().UpdateValue(variable, computable)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return false, nil
}

func (app *language) languageInstruction(
	testIns language_instruction.Instruction,
	baseDir string,
	index int,
	stackframe StackFrame,
	machine Machine,
	machineLanguage MachineLanguage,
) (bool, error) {
	if testIns.IsInstruction() {
		ins := testIns.Instruction()
		return app.instruction(ins, machine)
	}

	if testIns.IsCommand() {
		command := testIns.Command()
		err := machineLanguage.Command(command)
		return false, err
	}

	if testIns.IsMatch() {
		match := testIns.Match()
		err := machineLanguage.Match(match)
		return false, err
	}

	panic(errors.New("finish languageInstruction method in language, inside the interpreter"))
}

func (app *language) instruction(
	ins standard_instruction.Instruction,
	machine Machine,
) (bool, error) {
	err := machine.Receive(ins)
	if err != nil {
		return false, err
	}

	return false, nil
}
