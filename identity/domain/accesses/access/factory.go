package access

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
)

type factory struct {
	encPkFactory encryption.Factory
	sigPkFactory signature.PrivateKeyFactory
}

func createFactory(
	encPkFactory encryption.Factory,
	sigPkFactory signature.PrivateKeyFactory,
) Factory {
	out := factory{
		encPkFactory: encPkFactory,
		sigPkFactory: sigPkFactory,
	}

	return &out
}

// Create creates a new access instance
func (app *factory) Create() (Access, error) {
	id := uuid.NewV4()
	enc, err := app.encPkFactory.Create()
	if err != nil {
		return nil, err
	}

	sig := app.sigPkFactory.Create()
	return createAccess(&id, sig, enc), nil
}
