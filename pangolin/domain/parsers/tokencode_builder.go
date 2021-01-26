package parsers

import "errors"

type tokenCodeBuilder struct {
	content       VariableName
	tokenVariable string
}

func createTokenCodeBuilder() TokenCodeBuilder {
	out := tokenCodeBuilder{
		content:       nil,
		tokenVariable: "",
	}

	return &out
}

// Create initializes the builder
func (app *tokenCodeBuilder) Create() TokenCodeBuilder {
	return createTokenCodeBuilder()
}

// WithContent adds content to the builder
func (app *tokenCodeBuilder) WithContent(content VariableName) TokenCodeBuilder {
	app.content = content
	return app
}

// WithTokenVariable adds a tokenVariable to the builder
func (app *tokenCodeBuilder) WithTokenVariable(tokenVariable string) TokenCodeBuilder {
	app.tokenVariable = tokenVariable
	return app
}

// Now builds a new TokenCode instance
func (app *tokenCodeBuilder) Now() (TokenCode, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a TokenCode instance")
	}

	if app.tokenVariable == "" {
		return nil, errors.New("the tokenVariable is mandatory in order to build a TokenCode instance")
	}

	return createTokenCode(app.content, app.tokenVariable), nil
}
