package parsers

import "errors"

type tokenBuilder struct {
	codeMatch    CodeMatch
	tokenSection TokenSection
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		codeMatch:    nil,
		tokenSection: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithCodeMatch adds a codeMatch to the builder
func (app *tokenBuilder) WithCodeMatch(codeMatch CodeMatch) TokenBuilder {
	app.codeMatch = codeMatch
	return app
}

// WithTokenSection adds a tokenSection to the builder
func (app *tokenBuilder) WithTokenSection(tokenSection TokenSection) TokenBuilder {
	app.tokenSection = tokenSection
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.codeMatch != nil {
		return createTokenWithCodeMatch(app.codeMatch), nil
	}

	if app.tokenSection != nil {
		return createTokenWithTokenSection(app.tokenSection), nil
	}

	return nil, errors.New("the Token is invalid")
}
