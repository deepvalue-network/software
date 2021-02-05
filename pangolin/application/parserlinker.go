package scripts

import (
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type parserLinker struct {
	parser Parser
	linker Linker
}

func createParserLinker(
	parser Parser,
	linker Linker,
) ParserLinker {
	out := parserLinker{
		parser: parser,
		linker: linker,
	}

	return &out
}

// Execute executes a script from path
func (app *parserLinker) File(filePath string) (linkers.Program, error) {
	program, err := app.parser.File(filePath)
	if err != nil {
		return nil, err
	}

	return app.linker.Execute(program)
}

// Execute executes a script from string
func (app *parserLinker) Script(script string) (linkers.Program, error) {
	program, err := app.parser.Script(script)
	if err != nil {
		return nil, err
	}

	return app.linker.Execute(program)
}
