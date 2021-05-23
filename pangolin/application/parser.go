package application

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type parser struct {
	prs parsers.Parser
}

func createParser(
	prs parsers.Parser,
) Parser {
	out := parser{
		prs: prs,
	}

	return &out
}

// Execute executes a lexer and creates a parsed program
func (app *parser) Execute(lexer lexers.Lexer) (parsers.Program, error) {
	prog, err := app.prs.Execute(lexer)
	if err != nil {
		return nil, err
	}

	if castedProg, ok := prog.(parsers.Program); ok {
		return castedProg, nil
	}

	return nil, errors.New("the given lexer could not be parsed to a valid Program")
}
