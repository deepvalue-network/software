package identities

import (
	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	uuid "github.com/satori/go.uuid"
)

type connection struct {
	id      *uuid.UUID
	profile Profile
	conn    connections.Connection
	server  servers.Server
	sigPK   signature.PrivateKey
	encPK   encryption.PrivateKey
}

func createConnection(
	id *uuid.UUID,
	profile Profile,
	conn connections.Connection,
	server servers.Server,
	sigPK signature.PrivateKey,
	encPK encryption.PrivateKey,
) Connection {
	out := connection{
		id:      id,
		profile: profile,
		conn:    conn,
		server:  server,
		sigPK:   sigPK,
		encPK:   encPK,
	}

	return &out
}

// ID returns the id
func (obj *connection) ID() *uuid.UUID {
	return obj.id
}

// Profile returns the profile
func (obj *connection) Profile() Profile {
	return obj.profile
}

// Connection returns the connection
func (obj *connection) Connection() connections.Connection {
	return obj.conn
}

// Server returns the server
func (obj *connection) Server() servers.Server {
	return obj.server
}

// SigPK returns the signature PK
func (obj *connection) SigPK() signature.PrivateKey {
	return obj.sigPK
}

// EncPK returns the encryption PK
func (obj *connection) EncPK() encryption.PrivateKey {
	return obj.encPK
}
