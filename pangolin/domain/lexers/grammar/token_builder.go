package grammar

import (
	"errors"
)

type tokenBuilder struct {
	name    string
	grammar Grammar
	blocks  []TokenBlocks
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:    "",
		grammar: nil,
		blocks:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithBlocks adds a TokenBlocks to the builder
func (app *tokenBuilder) WithBlocks(blocks []TokenBlocks) TokenBuilder {
	app.blocks = blocks
	return app
}

// WithGrammar adds a Grammar to the builder
func (app *tokenBuilder) WithGrammar(grammar Grammar) TokenBuilder {
	app.grammar = grammar
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.blocks == nil {
		return nil, errors.New("the TokenBlocks is mandatory in order to build a Token instance")
	}

	if app.grammar != nil {
		return createTokenWithGrammar(delimiter, app.name, app.blocks, app.grammar), nil
	}

	return createToken(delimiter, app.name, app.blocks), nil
}
