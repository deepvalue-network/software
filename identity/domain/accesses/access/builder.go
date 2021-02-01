package access

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
)

type builder struct {
	id    *uuid.UUID
	sigPK signature.PrivateKey
	encPK encryption.PrivateKey
}

func createBuilder() Builder {
	out := builder{
		id:    nil,
		sigPK: nil,
		encPK: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithSignature adds a signature PK to the builder
func (app *builder) WithSignature(sigPK signature.PrivateKey) Builder {
	app.sigPK = sigPK
	return app
}

// WithEncryption adds a encryption PK to the builder
func (app *builder) WithEncryption(encPK encryption.PrivateKey) Builder {
	app.encPK = encPK
	return app
}

// Now builds a new Access instance
func (app *builder) Now() (Access, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build an Access instace")
	}

	if app.sigPK == nil {
		return nil, errors.New("the signature PK is mandatory in order to build an Access instance")
	}

	if app.encPK == nil {
		return nil, errors.New("the encryption PK is mandatory in order to build an Access instance")
	}

	return createAccess(app.id, app.sigPK, app.encPK), nil
}
