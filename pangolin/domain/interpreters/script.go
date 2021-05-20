package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type script struct {
	application       Application
	computableBuilder computable.Builder
	parser            parsers.Parser
	programBuilder    linkers.ProgramBuilder
	languageBuilder   linkers.LanguageBuilder
	linker            linkers.Linker
}

func createScript(
	application Application,
	computableBuilder computable.Builder,
	programBuilder linkers.ProgramBuilder,
	languageBuilder linkers.LanguageBuilder,
	linker linkers.Linker,
) Script {
	out := script{
		application:       application,
		computableBuilder: computableBuilder,
		programBuilder:    programBuilder,
		languageBuilder:   languageBuilder,
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
		linkedProg, err := app.linker.Execute(parsedProgram)
		if err != nil {
			return nil, err
		}

		if linkedProg.IsApplication() {
			return nil, errors.New("the linked program was expected to contain an application instance")
		}

		return linkedProg.Application(), nil
	}

	return nil, errors.New("the parsed compiled output was expected to contain a parsers.Program instance")
}
