package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type script struct {
	machineBuilder MachineBuilder
	valueBuilder   computable.Builder
	script         linkers.Script
}

func createScript(
	machineBuilder MachineBuilder,
	valueBuilder computable.Builder,
	linkedScript linkers.Script,
) Script {
	out := script{
		machineBuilder: machineBuilder,
		valueBuilder:   valueBuilder,
		script:         linkedScript,
	}

	return &out
}

// Execute executes a script in the interpreter
func (app *script) Execute(input map[string]computable.Value) (string, error) {
	langRef := app.script.Language()
	lang := langRef.Language()
	application := lang.Application()
	inVar := langRef.Input()
	code := app.script.Code()
	inStr, err := app.valueBuilder.Create().WithString(code).Now()
	if err != nil {
		return "", err
	}

	input[inVar] = inStr
	machine, err := app.machineBuilder.Create().WithInput(input).WithLanguage(lang).Now()
	if err != nil {
		return "", err
	}

	stackFrame, err := execute(machine, application)
	if err != nil {
		return "", err
	}

	outVar := langRef.Output()
	outVal, err := stackFrame.Current().Fetch(outVar)
	if err != nil {
		return "", err
	}

	if !outVal.IsString() {
		str := fmt.Sprintf("the output variable (%s) was expected to be a string", outVar)
		return "", errors.New(str)
	}

	return outVal.StringRepresentation(), nil
}
