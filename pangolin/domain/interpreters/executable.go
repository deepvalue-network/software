package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type executable struct {
	insMachineBuilder machines.InstructionBuilder
	stackFrameBuilder stackframes.Builder
	computableBuilder computable.Builder
	programBuilder    linkers.ProgramBuilder
	testableBuilder   linkers.TestableBuilder
	parser            parsers.Parser
	linker            linkers.Linker
}

func createExecutable(
	insMachineBuilder machines.InstructionBuilder,
	stackFrameBuilder stackframes.Builder,
	computableBuilder computable.Builder,
	programBuilder linkers.ProgramBuilder,
	testableBuilder linkers.TestableBuilder,
	parser parsers.Parser,
	linker linkers.Linker,
) Executable {
	out := executable{
		insMachineBuilder: insMachineBuilder,
		stackFrameBuilder: stackFrameBuilder,
		computableBuilder: computableBuilder,
		programBuilder:    programBuilder,
		testableBuilder:   testableBuilder,
		parser:            parser,
		linker:            linker,
	}

	return &out
}

// Execute executes the executable instance
func (app *executable) Execute(executable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error) {
	if executable.IsApplication() {
		appli := executable.Application()
		return app.Application(appli, input)
	}

	script := executable.Script()
	return app.Script(script, input)
}

// Application executes the application instance
func (app *executable) Application(appli linkers.Application, input stackframes.StackFrame) (stackframes.StackFrame, error) {
	labels := appli.Labels()
	machine, err := app.insMachineBuilder.Create().WithLabels(labels).WithStackFrame(input).Now()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	ins := appli.Instructions().All()
	for _, oneIns := range ins {
		err := machine.Receive(oneIns)
		if err != nil {
			return nil, err
		}
	}

	return input, nil
}

// Script executes the script instance
func (app *executable) Script(script linkers.Script, input stackframes.StackFrame) (stackframes.StackFrame, error) {
	langRef := script.Language()
	inVariable := langRef.Input()
	outVariable := script.Output()
	code := script.Code()

	codeValue, err := app.computableBuilder.Create().WithString(code).Now()
	if err != nil {
		return nil, err
	}

	scriptInputValues := map[string]computable.Value{
		inVariable: codeValue,
	}

	scriptInput := app.stackFrameBuilder.Create().WithVariables(scriptInputValues).Now()
	testable, err := app.testableBuilder.Create().WithLanguage(langRef).Now()
	if err != nil {
		return nil, err
	}

	linkedProgram, err := app.programBuilder.Create().WithTestable(testable).Now()
	if err != nil {
		return nil, err
	}

	if !linkedProgram.IsTestable() {
		return nil, errors.New("the linked program was expected to be a testable application")
	}

	linkedTestable := linkedProgram.Testable()
	if !linkedTestable.IsExecutable() {
		return nil, errors.New("the linked program was expected to be an executable application")
	}

	linkedExec := linkedTestable.Executable()
	retStackFrame, err := app.Execute(linkedExec, scriptInput)
	if err != nil {
		return nil, err
	}

	computedCodeValue, err := retStackFrame.Registry().Fetch(outVariable)
	if err != nil {
		return nil, err
	}

	if !computedCodeValue.IsString() {
		str := fmt.Sprintf("the output variable (%s) was expected to contain code and therefore be a string", outVariable)
		return nil, errors.New(str)
	}

	pOutputCode := computedCodeValue.String()
	programIns, err := app.parser.ExecuteScript(*pOutputCode)
	if err != nil {
		return nil, err
	}

	if parsedProgram, ok := programIns.(parsers.Program); ok {
		linkedProg, err := app.linker.Execute(parsedProgram)
		if err != nil {
			return nil, err
		}

		if linkedProg.IsTestable() {
			return nil, errors.New("the linked executable was expected to contain a testable program")
		}

		linkedTestable := linkedProg.Testable()
		if !linkedTestable.IsExecutable() {
			return nil, errors.New("the linked executable was expected to contain an executable program")
		}

		linkedExec := linkedTestable.Executable()
		return app.Execute(linkedExec, input)
	}

	return nil, errors.New("the parsed compiled output was expected to contain a parsers.Program instance")
}
