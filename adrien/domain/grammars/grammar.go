package grammars

import (
	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

type grammar struct {
	root     string
	tokens   tokens.Tokens
	rules    rules.Rules
	channels tokens.Tokens
}

func createGrammar(
	root string,
	tokens tokens.Tokens,
	rules rules.Rules,
) Grammar {
	return createGrammarInternally(root, tokens, rules, nil)
}

func createGrammarWithChannels(
	root string,
	tokens tokens.Tokens,
	rules rules.Rules,
	channels tokens.Tokens,
) Grammar {
	return createGrammarInternally(root, tokens, rules, channels)
}

func createGrammarInternally(
	root string,
	tokens tokens.Tokens,
	rules rules.Rules,
	channels tokens.Tokens,
) Grammar {
	out := grammar{
		root:     root,
		tokens:   tokens,
		channels: channels,
		rules:    rules,
	}

	return &out
}

// Root returns the root token name
func (obj *grammar) Root() string {
	return obj.root
}

// Tokens returns the tokens
func (obj *grammar) Tokens() tokens.Tokens {
	return obj.tokens
}

// Rules returns the rules
func (obj *grammar) Rules() rules.Rules {
	return obj.rules
}

// HasChannels returns true if there is a channel, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() tokens.Tokens {
	return obj.channels
}
