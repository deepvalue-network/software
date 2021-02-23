package connections

import (
	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	uuid "github.com/satori/go.uuid"
)

// Builder represenst a connection builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithProfile(profile Profile) Builder
	WithConnection(conn connections.Connection) Builder
	WithServer(server servers.Server) Builder
	WithSignaturePK(sigPK signature.PrivateKey) Builder
	WithEncryptionPK(encPK encryption.PrivateKey) Builder
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
