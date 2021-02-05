package scripts

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type parser struct {
	middleAdapter middle.Adapter
	parser        parsers.Parser
}

func createParser(
	middleAdapter middle.Adapter,
	prs parsers.Parser,
) Parser {
	out := parser{
		middleAdapter: middleAdapter,
		parser:        prs,
	}

	return &out
}

// File parse a file that contains a script and returns its program
func (app *parser) File(filePath string) (middle.Program, error) {
	ins, err := app.parser.ExecuteFile(filePath)
	if err != nil {
		return nil, err
	}

	if prog, ok := ins.(parsers.Program); ok {
		return app.middleAdapter.ToProgram(prog)
	}

	str := fmt.Sprintf("the file (%s) does not contain a valid program", filePath)
	return nil, errors.New(str)
}

// Script parse a script and returns its program
func (app *parser) Script(script string) (middle.Program, error) {
	ins, err := app.parser.ExecuteScript(script)
	if err != nil {
		return nil, err
	}

	if prog, ok := ins.(parsers.Program); ok {
		return app.middleAdapter.ToProgram(prog)
	}

	str := fmt.Sprintf("the given script does not contain a valid program, script:\n----\n%s", script)
	return nil, errors.New(str)
}
