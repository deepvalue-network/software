package parsers

import "errors"

type tokenSectionBuilder struct {
	variableName VariableName
	specific     SpecificTokenCode
}

func createTokenSectionBuilder() TokenSectionBuilder {
	out := tokenSectionBuilder{
		variableName: nil,
		specific:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenSectionBuilder) Create() TokenSectionBuilder {
	return createTokenSectionBuilder()
}

// WithVariableName adds a variableName to the builder
func (app *tokenSectionBuilder) WithVariableName(variableName VariableName) TokenSectionBuilder {
	app.variableName = variableName
	return app
}

// WithSpecific adds a specific code to the builder
func (app *tokenSectionBuilder) WithSpecific(specific SpecificTokenCode) TokenSectionBuilder {
	app.specific = specific
	return app
}

// Now builds a new TokenSection instance
func (app *tokenSectionBuilder) Now() (TokenSection, error) {
	if app.variableName != nil {
		return createTokenSectionWithVariableName(app.variableName), nil
	}

	if app.specific != nil {
		return createTokenSectionWithSpecific(app.specific), nil
	}

	return nil, errors.New("the TokenSection is invalid")
}
