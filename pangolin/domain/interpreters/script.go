package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type script struct {
	computableBuilder computable.Builder
	programBuilder    linkers.ProgramBuilder
	languageBuilder   linkers.LanguageBuilder
	application       Application
	language          Language
	parser            parsers.Parser
	linker            linkers.Linker
}

func createScript(
	computableBuilder computable.Builder,
	programBuilder linkers.ProgramBuilder,
	languageBuilder linkers.LanguageBuilder,
	application Application,
	language Language,
	parser parsers.Parser,
	linker linkers.Linker,
) Script {
	out := script{
		computableBuilder: computableBuilder,
		programBuilder:    programBuilder,
		languageBuilder:   languageBuilder,
		application:       application,
		language:          language,
		parser:            parser,
		linker:            linker,
	}

	return &out
}

// Execute converts a script to an application instance
func (app *script) Execute(script linkers.Script) (linkers.Application, error) {
	langRef := script.Language()
	langDef := langRef.Definition()
	langApp := langDef.Application()
	inVariable := langRef.Input()
	outVariable := script.Output()
	code := script.Code()

	codeValue, err := app.computableBuilder.Create().WithString(code).Now()
	if err != nil {
		return nil, err
	}

	input := map[string]computable.Value{
		inVariable: codeValue,
	}

	lang, err := app.languageBuilder.Create().WithApplication(langApp).Now()
	if err != nil {
		return nil, err
	}

	linkedProgram, err := app.programBuilder.Create().WithLanguage(lang).Now()
	if err != nil {
		return nil, err
	}

	if !linkedProgram.IsApplication() {
		return nil, errors.New("the linked program was expected to be an application")
	}

	linkedApp := linkedProgram.Application()
	stackFrame, err := app.application.Execute(linkedApp, input)
	if err != nil {
		return nil, err
	}

	computedCodeValue, err := stackFrame.Current().Fetch(outVariable)
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
		linkedExec, err := app.linker.Execute(parsedProgram)
		if err != nil {
			return nil, err
		}

		if linkedExec.IsApplication() {
			return nil, errors.New("the linked executable was expected to contain an application instance")
		}

		return linkedExec.Application(), nil
	}

	return nil, errors.New("the parsed compiled output was expected to contain a parsers.Program instance")
}

// Tests execute the script tests
func (app *script) Tests(script linkers.Script) error {
	// execute the language tests first:
	langDef := script.Language().Definition()
	err := app.language.Tests(langDef)
	if err != nil {
		return err
	}

	// if there is no tests, return successfully:
	if !script.HasTests() {
		return nil
	}

	// execute the script tests:
	tests := script.Tests()
	fmt.Printf(delimiter)
	fmt.Printf("Executing %d script tests...\n", len(tests))
	fmt.Printf(delimiter)
	for index, oneTest := range tests {
		name := oneTest.Name()
		script := oneTest.Script()

		fmt.Printf(delimiter)
		fmt.Printf(printTestStr, name)
		linkedApp, err := app.Execute(script)
		if err != nil {
			str := fmt.Sprintf("error while linking script to application... index: %d, error: %s", index, err.Error())
			return errors.New(str)
		}

		_, err = app.application.Execute(linkedApp, map[string]computable.Value{})
		if err != nil {
			str := fmt.Sprintf("error while executing script application... index: %d, error: %s", index, err.Error())
			return errors.New(str)
		}
	}

	fmt.Printf(delimiter)

	return nil
}
