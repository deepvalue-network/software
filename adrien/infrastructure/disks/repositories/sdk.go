package repositories

import (
	"github.com/deepvalue-network/software/adrien/domain/grammars"
	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

// NewGrammar creates a new disk grammar repository
func NewGrammar(
	fileRepository grammars.FileRepository,
) grammars.Repository {
	builder := grammars.NewBuilder()
	ruleAdapter := rules.NewAdapter()
	tokensAdapterBuilder := tokens.NewAdapterBuilder()
	return createGrammar(builder, fileRepository, ruleAdapter, tokensAdapterBuilder)
}

// NewGrammarFile creates a new disk grammar file repository
func NewGrammarFile(
	basePath string,
) grammars.FileRepository {
	fileBuilder := grammars.NewFileBuilder()
	return createGrammarFile(fileBuilder, basePath)
}
