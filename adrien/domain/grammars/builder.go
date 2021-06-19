package grammars

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

type builder struct {
	root     string
	tokens   tokens.Tokens
	rules    rules.Rules
	channels tokens.Tokens
}

func createBuilder() Builder {
	out := builder{
		root:     "",
		tokens:   nil,
		rules:    nil,
		channels: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens tokens.Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithRules add rules to the builder
func (app *builder) WithRules(rules rules.Rules) Builder {
	app.rules = rules
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels tokens.Tokens) Builder {
	app.channels = channels
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build a Grammar instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens are mandatory in order to build a Grammar instance")
	}

	if app.rules == nil {
		return nil, errors.New("the rules are mandatory in order to build a Grammar instance")
	}

	_, err := app.tokens.Find(app.root)
	if err != nil {
		str := fmt.Sprintf("the root Token instance (%s) could not be found in the Tokens instance", app.root)
		return nil, errors.New(str)
	}

	if app.channels != nil {
		return createGrammarWithChannels(app.root, app.tokens, app.rules, app.channels), nil
	}

	return createGrammar(app.root, app.tokens, app.rules), nil
}
