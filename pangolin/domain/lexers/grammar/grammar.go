package grammar

import (
	"errors"
	"fmt"
)

type grammar struct {
	name        string
	root        string
	rootToken   Token
	channels    Tokens
	tokens      Tokens
	rules       map[string]Rule
	subGrammars map[string]Grammar
}

func createGrammar(name string, root string, rootToken Token, rules map[string]Rule) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, nil, nil, nil)
}

func createGrammarWithSubGrammars(name string, root string, rootToken Token, rules map[string]Rule, subGrammars map[string]Grammar) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, nil, nil, subGrammars)
}

func createGrammarWithChannels(name string, root string, rootToken Token, rules map[string]Rule, channels Tokens) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, channels, nil, nil)
}

func createGrammarWithChannelsAndSubGrammars(name string, root string, rootToken Token, rules map[string]Rule, channels Tokens, subGrammars map[string]Grammar) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, channels, nil, subGrammars)
}

func createGrammarWithTokens(name string, root string, rootToken Token, rules map[string]Rule, tokens Tokens) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, nil, tokens, nil)
}

func createGrammarWithTokensAndSubGrammars(name string, root string, rootToken Token, rules map[string]Rule, tokens Tokens, subGrammars map[string]Grammar) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, nil, tokens, subGrammars)
}

func createGrammarWithChannelsAndTokens(name string, root string, rootToken Token, rules map[string]Rule, channels Tokens, tokens Tokens) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, channels, tokens, nil)
}

func createGrammarWithChannelsAndTokensAndSubGrammars(name string, root string, rootToken Token, rules map[string]Rule, channels Tokens, tokens Tokens, subGrammars map[string]Grammar) Grammar {
	return createGrammarInternally(name, root, rootToken, rules, channels, tokens, subGrammars)
}

func createGrammarInternally(
	name string,
	root string,
	rootToken Token,
	rules map[string]Rule,
	channels Tokens,
	tokens Tokens,
	subGrammars map[string]Grammar,
) Grammar {
	out := grammar{
		name:        name,
		root:        root,
		rootToken:   rootToken,
		rules:       rules,
		channels:    channels,
		tokens:      tokens,
		subGrammars: subGrammars,
	}

	return &out
}

// Name returns the name
func (obj *grammar) Name() string {
	return obj.name
}

// Root returns the root
func (obj *grammar) Root() string {
	return obj.root
}

// RootToken returns the rootToken
func (obj *grammar) RootToken() Token {
	return obj.rootToken
}

// Rules returns the rules
func (obj *grammar) Rules() map[string]Rule {
	return obj.rules
}

// HasTokens returns true if there is tokens, false otherwise
func (obj *grammar) HasTokens() bool {
	return obj.tokens != nil
}

// Tokens returns the tokens
func (obj *grammar) Tokens() Tokens {
	return obj.tokens
}

// HasSubGrammars returns true if there is sub grammars, false otherwise
func (obj *grammar) HasSubGrammars() bool {
	return obj.subGrammars != nil
}

// SubGrammars returns the sub grammars, if any
func (obj *grammar) SubGrammars() map[string]Grammar {
	return obj.subGrammars
}

// HasChannels returs true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels
func (obj *grammar) Channels() Tokens {
	return obj.channels
}

// FetchByName retrieves a grammar or a sub grammar by name
func (obj *grammar) FetchByName(name string) (Grammar, error) {
	if obj.name == name {
		return obj, nil
	}

	if !obj.HasSubGrammars() {
		str := fmt.Sprintf("the Grammar (%s) could not be found", name)
		return nil, errors.New(str)
	}

	return obj.subGrammars[name], nil
}
