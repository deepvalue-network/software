package parsers

import "errors"

type tokenSectionBuilder struct {
	code     TokenCode
	specific SpecificTokenCode
}

func createTokenSectionBuilder() TokenSectionBuilder {
	out := tokenSectionBuilder{
		code:     nil,
		specific: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSectionBuilder) Create() TokenSectionBuilder {
	return createTokenSectionBuilder()
}

// WithCode adds a code to the builder
func (app *tokenSectionBuilder) WithCode(code TokenCode) TokenSectionBuilder {
	app.code = code
	return app
}

// WithSpecific adds a specific code to the builder
func (app *tokenSectionBuilder) WithSpecific(specific SpecificTokenCode) TokenSectionBuilder {
	app.specific = specific
	return app
}

// Now builds a new TokenSection instance
func (app *tokenSectionBuilder) Now() (TokenSection, error) {
	if app.code != nil {
		return createTokenSectionWithCode(app.code), nil
	}

	if app.specific != nil {
		return createTokenSectionWithSpecific(app.specific), nil
	}

	return nil, errors.New("the TokenSection is invalid")
}
