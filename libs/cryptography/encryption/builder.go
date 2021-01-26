package encryption

import "errors"

type builder struct {
	password []byte
}

func createBuilder() Builder {
	out := builder{}
	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password []byte) Builder {
	app.password = password
	return app
}

// Now builds a new Encryption instance
func (app *builder) Now() (Encryption, error) {
	if app.password == nil {
		return nil, errors.New("the password is mandatory in order to build an Encryption instance")
	}

	hasher := curve.Hash()
	hasher.Write([]byte(app.password))
	hashedPass := hasher.Sum(nil)
	return createEncryption(hashedPass), nil
}
