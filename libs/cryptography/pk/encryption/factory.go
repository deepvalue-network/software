package encryption

import (
	"crypto/rand"
	"crypto/rsa"
)

type factory struct {
	builder Builder
	bitRate int
}

func createFactory(builder Builder, bitRate int) Factory {
	out := factory{
		builder: builder,
		bitRate: bitRate,
	}
	return &out
}

// Create generates a new PrivateKey instance
func (app *factory) Create() (PrivateKey, error) {
	pk, err := rsa.GenerateKey(rand.Reader, app.bitRate)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithPK(*pk).Now()
}
