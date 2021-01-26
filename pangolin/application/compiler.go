package scripts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/pangolin/domain/compilers"
	"github.com/steve-care-software/products/pangolin/domain/linkers"
)

type compiler struct {
	compilerApplication compilers.Application
	linker              Linker
}

func createCompiler(
	compilerApplication compilers.Application,
	linker Linker,
) Compiler {
	out := compiler{
		compilerApplication: compilerApplication,
		linker:              linker,
	}

	return &out
}

// Execute executes a compiler
func (app *compiler) Execute(script linkers.Script) (linkers.Application, error) {
	program, err := app.compilerApplication.Execute(script)
	if err != nil {
		return nil, err
	}

	linkedProgram, err := app.linker.Execute(program)
	if err != nil {
		return nil, err
	}

	if !linkedProgram.IsApplication() {
		str := fmt.Sprintf("the script was expected to compile to an application instance")
		return nil, errors.New(str)
	}

	return linkedProgram.Application(), nil
}
