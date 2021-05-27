package machines

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions/instruction"
)

type testInstruction struct {
	valueBuilder computable.Builder
	stackFrame   stackframes.StackFrame
	insApp       Instruction
	baseDir      string
}

func createTestInstruction(
	valueBuilder computable.Builder,
	stackFrame stackframes.StackFrame,
	insApp Instruction,
	baseDir string,
) TestInstruction {
	out := testInstruction{
		valueBuilder: valueBuilder,
		stackFrame:   stackFrame,
		insApp:       insApp,
		baseDir:      baseDir,
	}

	return &out
}

// Receive receives a test instruction
func (app *testInstruction) Receive(testIns test_instruction.Instruction) (bool, error) {
	if testIns.IsAssert() {
		assert := testIns.Assert()
		assertIndex := assert.Index()
		if assert.HasCondition() {
			condition := assert.Condition()
			condVal, err := app.stackFrame.Current().Fetch(condition)
			if err != nil {
				return false, err
			}

			if !condVal.IsBool() {
				return false, errors.New("the assert's condition was expected to contain a bool")
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
		err := app.insApp.Receive(ins)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	if testIns.IsReadFile() {
		readFile := testIns.ReadFile()
		relativePath := readFile.Path()
		joinedPath := filepath.Join(app.baseDir, relativePath)
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
		err = app.stackFrame.Current().UpdateValue(variable, computable)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return false, nil
}
