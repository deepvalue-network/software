package public

import (
	"crypto/rsa"
	"errors"
)

type builder struct {
	key *rsa.PublicKey
}

func createBuilder() Builder {
	out := builder{
		key: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithKey adds a key to the builder
func (app *builder) WithKey(key rsa.PublicKey) Builder {
	app.key = &key
	return app
}

// Now builds a new key instance
func (app *builder) Now() (Key, error) {
	if app.key == nil {
		return nil, errors.New("the rsa PublicKey is mandatory in order to build a PublicKey instance")
	}

	return createKey(*app.key), nil
}
