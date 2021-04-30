package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type script struct {
	machineLanguageBuilder MachineLanguageBuilder
	valueBuilder           computable.Builder
	script                 linkers.Script
}

func createScript(
	machineLanguageBuilder MachineLanguageBuilder,
	valueBuilder computable.Builder,
	linkedScript linkers.Script,
) Script {
	out := script{
		machineLanguageBuilder: machineLanguageBuilder,
		valueBuilder:           valueBuilder,
		script:                 linkedScript,
	}

	return &out
}

// Execute executes a script in the interpreter
func (app *script) Execute(input map[string]computable.Value) (string, error) {
	code := app.script.Code()
	inStr, err := app.valueBuilder.Create().WithString(code).Now()
	if err != nil {
		return "", err
	}

	langRef := app.script.Language()
	langDef := langRef.Definition()
	inVar := langRef.Input()
	input[inVar] = inStr
	machine, err := app.machineLanguageBuilder.Create().WithInput(input).WithLanguage(langDef).Now()
	if err != nil {
		return "", err
	}

	application := langDef.Application()
	stackFrame, err := app.execute(machine, application)
	if err != nil {
		return "", err
	}

	outVar := app.script.Output()
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

func (app *script) execute(machine MachineLanguage, linkedLangApp linkers.LanguageApplication) (StackFrame, error) {
	ins := linkedLangApp.Instructions().All()
	for _, oneIns := range ins {
		err := machine.Receive(oneIns)
		if err != nil {
			return nil, err
		}
	}

	return machine.StackFrame(), nil
}
