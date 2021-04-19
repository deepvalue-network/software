package parsers

import "errors"

type languageValueBuilder struct {
	root           string
	tokens         RelativePath
	channels       RelativePath
	rules          RelativePath
	logic          RelativePath
	input          string
	extends        []RelativePath
	patternMatches []PatternMatch
}

func createLanguageValueBuilder() LanguageValueBuilder {
	out := languageValueBuilder{
		root:           "",
		tokens:         nil,
		channels:       nil,
		rules:          nil,
		logic:          nil,
		input:          "",
		extends:        nil,
		patternMatches: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageValueBuilder) Create() LanguageValueBuilder {
	return createLanguageValueBuilder()
}

// WithRoot add root to the builder
func (app *languageValueBuilder) WithRoot(root string) LanguageValueBuilder {
	app.root = root
	return app
}

// WithTokens add tokens to the builder
func (app *languageValueBuilder) WithTokens(tokens RelativePath) LanguageValueBuilder {
	app.tokens = tokens
	return app
}

// WithChannels add channels to the builder
func (app *languageValueBuilder) WithChannels(channels RelativePath) LanguageValueBuilder {
	app.channels = channels
	return app
}

// WithRules add rules to the builder
func (app *languageValueBuilder) WithRules(rules RelativePath) LanguageValueBuilder {
	app.rules = rules
	return app
}

// WithLogic adds logic to the builder
func (app *languageValueBuilder) WithLogic(logic RelativePath) LanguageValueBuilder {
	app.logic = logic
	return app
}

// WithInputVariable adds an input variable to the builder
func (app *languageValueBuilder) WithInputVariable(inputVar string) LanguageValueBuilder {
	app.input = inputVar
	return app
}

// WithExtends add extends variable to the builder
func (app *languageValueBuilder) WithExtends(extends []RelativePath) LanguageValueBuilder {
	app.extends = extends
	return app
}

// WithPatternMatches add pattern matches to the builder
func (app *languageValueBuilder) WithPatternMatches(patternMatches []PatternMatch) LanguageValueBuilder {
	app.patternMatches = patternMatches
	return app
}

// Now builds a new LanguageValue instance
func (app *languageValueBuilder) Now() (LanguageValue, error) {
	if app.root != "" {
		return createLanguageValueWithRoot(app.root), nil
	}

	if app.tokens != nil {
		return createLanguageValueWithTokens(app.tokens), nil
	}

	if app.channels != nil {
		return createLanguageValueWithChannels(app.channels), nil
	}

	if app.rules != nil {
		return createLanguageValueWithRules(app.rules), nil
	}

	if app.logic != nil {
		return createLanguageValueWithLogic(app.logic), nil
	}

	if app.input != "" {
		return createLanguageValueWithInput(app.input), nil
	}

	if app.extends != nil {
		return createLanguageValueWithExtends(app.extends), nil
	}

	if app.patternMatches != nil {
		return createLanguageValueWithPatternMatches(app.patternMatches), nil
	}

	return nil, errors.New("the LanguageValue is invalid")
}
