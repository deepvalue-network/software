package lexer

import (
	"github.com/deepvalue-network/software/adrien/domain/asts"
	"github.com/deepvalue-network/software/adrien/domain/grammars"
)

// Application represents a lexer application
type Application interface {
	Execute(script string, grammar grammars.Grammar) (asts.AST, error)
}
