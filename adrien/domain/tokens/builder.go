package tokens

import "errors"

type builder struct {
	list []Token
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens []Token) Builder {
	app.list = tokens
	return app
}

// Now builds a new Tokens instance
func (app *builder) Now() (Tokens, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Tokens instance")
	}

	mp := map[string]Token{}
	for _, oneToken := range app.list {
		name := oneToken.Name()
		mp[name] = oneToken
	}

	return createTokens(app.list, mp), nil
}
