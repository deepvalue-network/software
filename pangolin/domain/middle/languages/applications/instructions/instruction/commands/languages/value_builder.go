package languages

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type valueBuilder struct {
	root           string
	tokensPath     string
	rulesPath      string
	logicsPath     string
	patternMatches []definitions.PatternMatch
	inputVariable  string
	channelsPath   string
	extends        []string
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		root:           "",
		tokensPath:     "",
		rulesPath:      "",
		logicsPath:     "",
		patternMatches: nil,
		inputVariable:  "",
		channelsPath:   "",
		extends:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithRoot adds a root to the builder
func (app *valueBuilder) WithRoot(root string) ValueBuilder {
	app.root = root
	return app
}

// WithTokensPath adds a tokensPath to the builder
func (app *valueBuilder) WithTokensPath(tokensPath string) ValueBuilder {
	app.tokensPath = tokensPath
	return app
}

// WithRulesPath adds a rulesPath to the builder
func (app *valueBuilder) WithRulesPath(rulesPath string) ValueBuilder {
	app.rulesPath = rulesPath
	return app
}

// WithLogicsPath adds a logicsPath to the builder
func (app *valueBuilder) WithLogicsPath(logicsPath string) ValueBuilder {
	app.logicsPath = logicsPath
	return app
}

// WithPatternMatches adds patternMatches to the builder
func (app *valueBuilder) WithPatternMatches(patternMatches []definitions.PatternMatch) ValueBuilder {
	app.patternMatches = patternMatches
	return app
}

// WithInputVariable adds an inputVariable to the builder
func (app *valueBuilder) WithInputVariable(inputVariable string) ValueBuilder {
	app.inputVariable = inputVariable
	return app
}

// WithChannelsPath adds a channelsPath to the builder
func (app *valueBuilder) WithChannelsPath(channelsPath string) ValueBuilder {
	app.channelsPath = channelsPath
	return app
}

// WithExtends adds an extends to the builder
func (app *valueBuilder) WithExtends(extends []string) ValueBuilder {
	app.extends = extends
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.root != "" {
		return createValueWithRoot(app.root), nil
	}

	if app.tokensPath != "" {
		return createValueWithTokensPath(app.tokensPath), nil
	}

	if app.rulesPath != "" {
		return createValueWithRulesPath(app.rulesPath), nil
	}

	if app.logicsPath != "" {
		return createValueWithLogicsPath(app.logicsPath), nil
	}

	if app.patternMatches != nil {
		return createValueWithPatternMatches(app.patternMatches), nil
	}

	if app.inputVariable != "" {
		return createValueWithInputVariable(app.inputVariable), nil
	}

	if app.channelsPath != "" {
		return createValueWithChannelsPath(app.channelsPath), nil
	}

	if app.extends != nil {
		return createValueWithExtends(app.extends), nil
	}

	return nil, errors.New("the Value is invalid")
}
