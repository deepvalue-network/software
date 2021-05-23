package application

import "github.com/deepvalue-network/software/pangolin/domain/lexers"

type lexer struct {
	adapter lexers.Adapter
}

func createLexer(
	adapter lexers.Adapter,
) Lexer {
	out := lexer{
		adapter: adapter,
	}

	return &out
}

// Execute executes the lexer
func (app *lexer) Execute(script string) (lexers.Lexer, error) {
	return app.adapter.ToLexer(script)
}
