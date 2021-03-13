package middle

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/targets"
)

type languageBuilder struct {
	root           string
	tokens         string
	rules          string
	logics         string
	patternMatches []PatternMatch
	input          string
	output         string
	targets        targets.Targets
	channels       string
	extends        []string
}

func createLanguageBuilder() LanguageBuilder {
	out := languageBuilder{
		root:           "",
		tokens:         "",
		rules:          "",
		logics:         "",
		patternMatches: nil,
		input:          "",
		output:         "",
		targets:        nil,
		channels:       "",
		extends:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder()
}

// WithRoot adds a root pattern to the builder
func (app *languageBuilder) WithRoot(root string) LanguageBuilder {
	app.root = root
	return app
}

// WithTokensPath adds a token path to the builder
func (app *languageBuilder) WithTokensPath(tokens string) LanguageBuilder {
	app.tokens = tokens
	return app
}

// WithChannelsPath adds a channels path to the builder
func (app *languageBuilder) WithChannelsPath(channels string) LanguageBuilder {
	app.channels = channels
	return app
}

// WithRulesPath adds a rules path to the builder
func (app *languageBuilder) WithRulesPath(rules string) LanguageBuilder {
	app.rules = rules
	return app
}

// WithLogicsPath adds a logics path to the builder
func (app *languageBuilder) WithLogicsPath(logics string) LanguageBuilder {
	app.logics = logics
	return app
}

// WithPatternMatches add pattern matches to the builder
func (app *languageBuilder) WithPatternMatches(patternMatches []PatternMatch) LanguageBuilder {
	app.patternMatches = patternMatches
	return app
}

// WithInputVariable adds an inputVariable to the builder
func (app *languageBuilder) WithInputVariable(input string) LanguageBuilder {
	app.input = input
	return app
}

// WithOutputVariable adds an output variable to the builder
func (app *languageBuilder) WithOutputVariable(output string) LanguageBuilder {
	app.output = output
	return app
}

// WithTargets add targets to the builder
func (app *languageBuilder) WithTargets(targets targets.Targets) LanguageBuilder {
	app.targets = targets
	return app
}

// WithExtends add extends bucketPaths to the builder
func (app *languageBuilder) WithExtends(extends []string) LanguageBuilder {
	app.extends = extends
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
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

	if app.output == "" {
		return nil, errors.New("the output variable is mandatory in order to build a Language instance")
	}

	if app.targets == nil {
		return nil, errors.New("the targets are mandatory in order to build a Language instance")
	}

	if app.channels != "" && app.extends != nil {
		return createLanguageWithChannelsAndExtends(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.output, app.targets, app.channels, app.extends), nil
	}

	if app.channels != "" {
		return createLanguageWithChannels(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.output, app.targets, app.channels), nil
	}

	if app.extends != nil {
		return createLanguageWithExtends(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.output, app.targets, app.extends), nil
	}

	return createLanguage(app.root, app.tokens, app.rules, app.logics, app.patternMatches, app.input, app.output, app.targets), nil
}
