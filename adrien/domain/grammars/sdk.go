package grammars

import (
	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewFileBuilder creates a new file builder
func NewFileBuilder() FileBuilder {
	return createFileBuilder()
}

// Repository represents a grammar repository
type Repository interface {
	Retrieve(relativePath string) (Grammar, error)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root string) Builder
	WithTokens(tokens tokens.Tokens) Builder
	WithRules(rules rules.Rules) Builder
	WithChannels(channels tokens.Tokens) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Root() string
	Tokens() tokens.Tokens
	Rules() rules.Rules
	HasChannels() bool
	Channels() tokens.Tokens
}

// FileRepository represents a file repository
type FileRepository interface {
	Retrieve(relativePath string) (File, error)
}

// FileBuilder represents a file builder
type FileBuilder interface {
	Create() FileBuilder
	WithRoot(root string) FileBuilder
	WithTokensPath(tokensPath string) FileBuilder
	WithRulesPath(rulesPath string) FileBuilder
	WithChannelsPath(channelsPath string) FileBuilder
	Now() (File, error)
}

// File represents a grammar file
type File interface {
	Root() string
	TokensPath() string
	RulesPath() string
	HasChannelsPath() bool
	ChannelsPath() string
}
