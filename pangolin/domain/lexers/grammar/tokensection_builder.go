package grammar

import "errors"

type tokenSectionBuilder struct {
	rule  TokenRule
	token RawToken
}

func createTokenSectionBuilder() TokenSectionBuilder {
	out := tokenSectionBuilder{
		rule:  nil,
		token: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSectionBuilder) Create() TokenSectionBuilder {
	return createTokenSectionBuilder()
}

// WithRule add a rule to the builder
func (app *tokenSectionBuilder) WithRule(rule TokenRule) TokenSectionBuilder {
	app.rule = rule
	return app
}

// WithToken add a tokenBlock to the builder
func (app *tokenSectionBuilder) WithToken(token RawToken) TokenSectionBuilder {
	app.token = token
	return app
}

// Now builds a new TokenSection instance
func (app *tokenSectionBuilder) Now() (TokenSection, error) {
	if app.token != nil {
		return createTokenSectionWithToken(app.token), nil
	}

	if app.rule != nil {
		return createTokenSectionWithRule(app.rule), nil
	}

	return nil, errors.New("the TokenSection is invalid")
}
