package grammar

import "errors"

type replacementTokenBuilder struct {
	toGrammar string
	fromToken string
}

func createReplacementTokenBuilder() ReplacementTokenBuilder {
	out := replacementTokenBuilder{
		toGrammar: "",
		fromToken: "",
	}

	return &out
}

// Create initializes the builder
func (app *replacementTokenBuilder) Create() ReplacementTokenBuilder {
	return createReplacementTokenBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *replacementTokenBuilder) WithToGrammar(toGrammar string) ReplacementTokenBuilder {
	app.toGrammar = toGrammar
	return app
}

// WithFromToken adds a from token to the builder
func (app *replacementTokenBuilder) WithFromToken(fromToken string) ReplacementTokenBuilder {
	app.fromToken = fromToken
	return app
}

// Now builds a new ReplacementToken instance
func (app *replacementTokenBuilder) Now() (ReplacementToken, error) {
	if app.toGrammar == "" {
		return nil, errors.New("the to Grammar is mandatory in order to build a ReplacementToken instance")
	}

	if app.fromToken == "" {
		return nil, errors.New("the from Token is mandatory in order to build a ReplacementToken instance")
	}

	return createReplacementToken(app.toGrammar, app.fromToken)
}
