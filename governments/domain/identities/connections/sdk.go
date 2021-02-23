package connections

import (
	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	uuid "github.com/satori/go.uuid"
)

// Builder represents a connections builder
type Builder interface {
	Create() Builder
	WithConnections(connections []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	All() []Connection
}

// ConnectionBuilder represenst a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithID(id *uuid.UUID) ConnectionBuilder
	WithProfile(profile Profile) ConnectionBuilder
	WithConnection(conn connections.Connection) ConnectionBuilder
	WithServer(server servers.Server) ConnectionBuilder
	WithSignaturePK(sigPK signature.PrivateKey) ConnectionBuilder
	WithEncryptionPK(encPK encryption.PrivateKey) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	ID() *uuid.UUID
	Profile() Profile
	Connection() connections.Connection
	Server() servers.Server
	SigPK() signature.PrivateKey
	EncPK() encryption.PrivateKey
}

// ProfileBuilder represents a profile builder
type ProfileBuilder interface {
	Create() ProfileBuilder
	WithName(name string) ProfileBuilder
	WithRank(rank uint) ProfileBuilder
	Now() (Profile, error)
}

// Profile represents a connection profile
type Profile interface {
	Name() string
	HasRank() bool
	Rank() uint
}
