package connections

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	uuid "github.com/satori/go.uuid"
)

type connectionBuilder struct {
	sigPKFactory signature.PrivateKeyFactory
	encPKFactory encryption.Factory
	id           *uuid.UUID
	profile      Profile
	conn         connections.Connection
	server       servers.Server
	sigPK        signature.PrivateKey
	encPK        encryption.PrivateKey
}

func createConnectionBuilder(
	sigPKFactory signature.PrivateKeyFactory,
	encPKFactory encryption.Factory,
) ConnectionBuilder {
	out := connectionBuilder{
		sigPKFactory: sigPKFactory,
		encPKFactory: encPKFactory,
		id:           nil,
		profile:      nil,
		conn:         nil,
		server:       nil,
		sigPK:        nil,
		encPK:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder(
		app.sigPKFactory,
		app.encPKFactory,
	)
}

// WithID adds an ID to the builder
func (app *connectionBuilder) WithID(id *uuid.UUID) ConnectionBuilder {
	app.id = id
	return app
}

// WithProfile adds a profile to the builder
func (app *connectionBuilder) WithProfile(profile Profile) ConnectionBuilder {
	app.profile = profile
	return app
}

// WithConnection adds a connection to the builder
func (app *connectionBuilder) WithConnection(conn connections.Connection) ConnectionBuilder {
	app.conn = conn
	return app
}

// WithServer adds a server to the builder
func (app *connectionBuilder) WithServer(server servers.Server) ConnectionBuilder {
	app.server = server
	return app
}

// WithSignaturePK adds a signaturePK to the builder
func (app *connectionBuilder) WithSignaturePK(sigPK signature.PrivateKey) ConnectionBuilder {
	app.sigPK = sigPK
	return app
}

// WithEncryptionPK adds an encryption to the builder
func (app *connectionBuilder) WithEncryptionPK(encPK encryption.PrivateKey) ConnectionBuilder {
	app.encPK = encPK
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.profile == nil {
		return nil, errors.New("the profile is mandatory in order to build an identity Connection instance")
	}

	if app.conn == nil {
		return nil, errors.New("the connection is mandatory in order to build an identity Connection instance")
	}

	if app.server == nil {
		return nil, errors.New("the server is mandatory in order to build an identity Connection instance")
	}

	if app.sigPK == nil {
		app.sigPK = app.sigPKFactory.Create()
	}

	if app.encPK == nil {
		encPK, err := app.encPKFactory.Create()
		if err != nil {
			return nil, err
		}

		app.encPK = encPK
	}

	if app.id == nil {
		id := uuid.NewV4()
		app.id = &id
	}

	return createConnection(app.id, app.profile, app.conn, app.server, app.sigPK, app.encPK), nil
}
