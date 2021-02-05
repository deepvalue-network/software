package access

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
)

type factory struct {
	encPkFactory encryption.Factory
	sigPkFactory signature.PrivateKeyFactory
	builder      Builder
}

func createFactory(
	encPkFactory encryption.Factory,
	sigPkFactory signature.PrivateKeyFactory,
	builder Builder,
) Factory {
	out := factory{
		encPkFactory: encPkFactory,
		sigPkFactory: sigPkFactory,
		builder:      builder,
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
	return app.builder.Create().WithID(&id).WithEncryption(enc).WithSignature(sig).Now()
}
