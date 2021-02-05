package resources

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
)

type accessBuilder struct {
	resource  Mutable
	owners    []*uuid.UUID
	encrypted public.Key
}

func createAccessBuilder() AccessBuilder {
	out := accessBuilder{
		resource:  nil,
		owners:    nil,
		encrypted: nil,
	}

	return &out
}

// Create initializes the builder
func (app *accessBuilder) Create() AccessBuilder {
	return createAccessBuilder()
}

// WithResource adds a resource to the builder
func (app *accessBuilder) WithResource(res Mutable) AccessBuilder {
	app.resource = res
	return app
}

// WithOwners adds owners to the builder
func (app *accessBuilder) WithOwners(owners []*uuid.UUID) AccessBuilder {
	app.owners = owners
	return app
}

// WithEncryptionPubKey adds the encryption pubKey to the builder
func (app *accessBuilder) WithEncryptionPubKey(pubKey public.Key) AccessBuilder {
	app.encrypted = pubKey
	return app
}

// Now builds a new Access instance
func (app *accessBuilder) Now() (Access, error) {
	if app.owners != nil && len(app.owners) <= 0 {
		app.owners = nil
	}

	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Access instance")
	}

	if app.owners == nil {
		return nil, errors.New("the owners are mandatory in order to build an Access instance")
	}

	if app.encrypted != nil {
		return createAccessWithEncryptedPubkey(app.resource, app.owners, app.encrypted), nil
	}

	return createAccess(app.resource, app.owners), nil
}
