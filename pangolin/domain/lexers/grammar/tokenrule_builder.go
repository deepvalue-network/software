package grammar

import "errors"

type tokenRuleBuilder struct {
	rule     Rule
	rawToken RawToken
}

func createTokenRuleBuilder() TokenRuleBuilder {
	out := tokenRuleBuilder{
		rule:     nil,
		rawToken: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenRuleBuilder) Create() TokenRuleBuilder {
	return createTokenRuleBuilder()
}

// WithRule adds a rule to the builder
func (app *tokenRuleBuilder) WithRule(rule Rule) TokenRuleBuilder {
	app.rule = rule
	return app
}

// WithRawToken adds a rawToken to the builder
func (app *tokenRuleBuilder) WithRawToken(rawToken RawToken) TokenRuleBuilder {
	app.rawToken = rawToken
	return app
}

// Now builds a new TokenRule instance
func (app *tokenRuleBuilder) Now() (TokenRule, error) {
	if app.rule == nil {
		return nil, errors.New("the Rule is mandatory in order to build a TokenRule instance")
	}

	if app.rawToken == nil {
		return nil, errors.New("the RawToken is mandatory in order to build a TokenRule instance")
	}

	return createTokenRule(app.rule, app.rawToken), nil
}
