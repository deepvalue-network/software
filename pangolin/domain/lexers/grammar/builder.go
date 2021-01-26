package grammar

import (
	"errors"
	"fmt"
)

type builder struct {
	name        string
	root        string
	channels    Tokens
	tokens      Tokens
	rules       []Rule
	subGrammars map[string]Grammar
	gr          Grammar
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		root:        "",
		channels:    nil,
		tokens:      nil,
		rules:       nil,
		subGrammars: nil,
		gr:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName add a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithRoot add root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels Tokens) Builder {
	app.channels = channels
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithRules add rules to the builder
func (app *builder) WithRules(rules []Rule) Builder {
	app.rules = rules
	return app
}

// WithSubGrammars add sub grammars to the builder
func (app *builder) WithSubGrammars(grammars map[string]Grammar) Builder {
	app.subGrammars = grammars
	return app
}

// WithGrammar add a grammar to the builder
func (app *builder) WithGrammar(gr Grammar) Builder {
	app.gr = gr
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Grammar instance")
	}

	var channels Tokens
	var tokens Tokens
	rules := map[string]Rule{}
	subGrammars := map[string]Grammar{}
	if app.gr != nil {
		app.root = app.gr.Root()
		rules = app.gr.Rules()
		channels = app.gr.Channels()
		tokens = app.gr.Tokens()
		if app.gr.HasSubGrammars() {
			subGrammars = app.gr.SubGrammars()
		}
	} else {
		if app.root == "" {
			return nil, errors.New("the root Token is mandatory in order to build a Grammar instance")
		}

		if app.rules == nil {
			return nil, errors.New("the []Rule are mandatory in order to build a Grammar instance")
		}

		for _, oneRule := range app.rules {
			rules[oneRule.Name()] = oneRule
		}

		if app.channels != nil {
			channels = app.channels
		}

		if app.tokens != nil {
			tokens = app.tokens
		}

		if app.subGrammars != nil {
			subGrammars = app.subGrammars
		}
	}

	if tokens == nil {
		return nil, errors.New("the Tokens instance is mandatory in order to build a Grammar instance")
	}

	allTokens := tokens.Tokens()
	if rootToken, ok := allTokens[app.root]; ok {
		if (channels != nil) && (tokens != nil) && len(subGrammars) > 0 {
			return createGrammarWithChannelsAndTokensAndSubGrammars(app.name, app.root, rootToken, rules, channels, tokens, subGrammars), nil
		}

		if (channels != nil) && (tokens != nil) {
			return createGrammarWithChannelsAndTokens(app.name, app.root, rootToken, rules, channels, tokens), nil
		}

		if (channels != nil) && len(subGrammars) > 0 {
			return createGrammarWithChannelsAndSubGrammars(app.name, app.root, rootToken, rules, channels, subGrammars), nil
		}

		if channels != nil {
			return createGrammarWithChannels(app.name, app.root, rootToken, rules, channels), nil
		}

		if tokens != nil {
			return createGrammarWithTokens(app.name, app.root, rootToken, rules, tokens), nil
		}

		if (tokens != nil) && len(subGrammars) > 0 {
			return createGrammarWithTokensAndSubGrammars(app.name, app.root, rootToken, rules, tokens, subGrammars), nil
		}

		if len(subGrammars) > 0 {
			return createGrammarWithSubGrammars(app.name, app.root, rootToken, rules, subGrammars), nil
		}

		return createGrammar(app.name, app.root, rootToken, rules), nil
	}

	str := fmt.Sprintf("root (name: %s) is not a declared token", app.root)
	return nil, errors.New(str)

}
