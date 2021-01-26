package encryption

import (
	"crypto/rsa"
	"errors"

	"github.com/steve-care-software/products/pangolin/libs/cryptography/pk/encryption/public"
)

type builder struct {
	pubKeyBuilder public.Builder
	pk            *rsa.PrivateKey
}

func createBuilder(pubKeyBuilder public.Builder) Builder {
	out := builder{
		pubKeyBuilder: pubKeyBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.pubKeyBuilder)
}

// WithPK adds a privateKey to the builder
func (app *builder) WithPK(pk rsa.PrivateKey) Builder {
	app.pk = &pk
	return app
}

// Now builds a new PrivateKey instance
func (app *builder) Now() (PrivateKey, error) {
	if app.pk == nil {
		return nil, errors.New("the rsa PrivateKey is mandatory in order to build an encryption PrivateKey instance")
	}

	pubKey, err := app.pubKeyBuilder.Create().WithKey(app.pk.PublicKey).Now()
	if err != nil {
		return nil, err
	}

	return createPrivateKey(*app.pk, pubKey), nil
}
