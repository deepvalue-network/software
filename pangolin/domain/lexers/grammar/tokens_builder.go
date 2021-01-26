package grammar

import "errors"

type tokensBuilder struct {
	tokens       map[string]Token
	replacements []ReplacementToken
}

func createTokensBuilder() TokensBuilder {
	out := tokensBuilder{
		tokens:       nil,
		replacements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokensBuilder) Create() TokensBuilder {
	return createTokensBuilder()
}

// WithTokens add tokens to the builder
func (app *tokensBuilder) WithTokens(tokens map[string]Token) TokensBuilder {
	app.tokens = tokens
	return app
}

// WithReplacements add replacement tokens to the builder
func (app *tokensBuilder) WithReplacements(replacements []ReplacementToken) TokensBuilder {
	app.replacements = replacements
	return app
}

// Now build a new Tokens instance
func (app *tokensBuilder) Now() (Tokens, error) {
	if app.tokens == nil {
		return nil, errors.New("the Token map is mandatory in order to build a Tokens instance")
	}

	if app.replacements != nil {
		return createTokensWithReplacements(app.tokens, app.replacements)
	}

	return createTokens(app.tokens)
}
