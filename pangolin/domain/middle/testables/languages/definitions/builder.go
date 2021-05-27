package definitions

import (
	"errors"
)

type builder struct {
	root           string
	tokens         string
	rules          string
	logics         string
	patternMatches []PatternMatch
	input          string
	channels       string
	extends        []string
}

func createBuilder() Builder {
	out := builder{
		root:           "",
		tokens:         "",
		rules:          "",
		logics:         "",
		patternMatches: nil,
		input:          "",
		channels:       "",
		extends:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root pattern to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// WithTokensPath adds a token path to the builder
func (app *builder) WithTokensPath(tokens string) Builder {
	app.tokens = tokens
	return app
}

// WithChannelsPath adds a channels path to the builder
func (app *builder) WithChannelsPath(channels string) Builder {
	app.channels = channels
	return app
}

// WithRulesPath adds a rules path to the builder
func (app *builder) WithRulesPath(rules string) Builder {
	app.rules = rules
	return app
}

// WithLogicsPath adds a logics path to the builder
func (app *builder) WithLogicsPath(logics string) Builder {
	app.logics = logics
	return app
}

// WithPatternMatches add pattern matches to the builder
func (app *builder) WithPatternMatches(patternMatches []PatternMatch) Builder {
	app.patternMatches = patternMatches
	return app
}

// WithInputVariable adds an inputVariable to the builder
func (app *builder) WithInputVariable(input string) Builder {
	app.input = input
	return app
}

// WithExtends add extends bucketPaths to the builder
func (app *builder) WithExtends(extends []string) Builder {
	app.extends = extends
	return app
}

// Now builds a new Language instance
func (app *builder) Now() (Definition, error) {
	if app.tokens == "" {
		return nil, errors.New("the tokens path is mandatory in order to build a Language instance")
	}

	if app.root == "" {
		return nil, errors.New("the root pattern is mandatory in order to build a Language instance")
	}

	if app.rules == "" {
		return nil, errors.New("the rules path is mandatory in order to build a Language instance")
	}

	if app.logics == "" {
		return nil, errors.New("the logics path is mandatory in order to build a Language instance")
	}

	if len(app.patternMatches) <= 0 {
		app.patternMatches = nil
	}

	if app.patternMatches == nil {
		return nil, errors.New("the []PatternMatch are mandatory in order to build a Language instance")
	}

	if app.input == "" {
		return nil, errors.New("the input variable is mandatory in order to build a Language instance")
	}

	if app.channels != "" && app.extends != nil {
		return createDefinitionWithChannelsAndExtends(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.channels, app.extends), nil
	}

	if app.channels != "" {
		return createDefinitionWithChannels(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.channels), nil
	}

	if app.extends != nil {
		return createDefinitionWithExtends(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.extends), nil
	}

	return createDefinition(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input), nil
}
