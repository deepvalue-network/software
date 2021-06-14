package grammars

import (
	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

// Repository represents a grammar repository
type Repository interface {
	Retrieve(relativePath string) (Grammar, error)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root string) Builder
	WithTokens(tokens tokens.Tokens) Builder
	WithChannels(channels tokens.Tokens) Builder
	WithRules(rules rules.Rules) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Root() string
	Tokens() tokens.Tokens
	Channels() tokens.Tokens
	Rules() rules.Rules
}
