package token

import "errors"

type builder struct {
	codeMatch CodeMatch
	code      Code
}

func createBuilder() Builder {
	out := builder{
		codeMatch: nil,
		code:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCodeMatch adds a codeMatch to the builder
func (app *builder) WithCodeMatch(codeMatch CodeMatch) Builder {
	app.codeMatch = codeMatch
	return app
}

// WithCode adds a code to the builder
func (app *builder) WithCode(code Code) Builder {
	app.code = code
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.codeMatch != nil {
		return createTokenWithCodeMatch(app.codeMatch), nil
	}

	if app.code != nil {
		return createTokenWithCode(app.code), nil
	}

	return nil, errors.New("the Token is invalid")
}
