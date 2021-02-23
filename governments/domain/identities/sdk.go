package identities

import (
	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	connectionsBuilder := NewConnectionsBuilder()
	return createBuilder(connectionsBuilder)
}

// NewConnectionsBuilder creates a new connections builder
func NewConnectionsBuilder() ConnectionsBuilder {
	return createConnectionsBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder(encBitrate int) ConnectionBuilder {
	sigPKFactory := signature.NewPrivateKeyFactory()
	encPKFactory := encryption.NewFactory(encBitrate)
	return createConnectionBuilder(sigPKFactory, encPKFactory)
}

// NewProfileBuilder creates a new profile builder
func NewProfileBuilder() ProfileBuilder {
	return createProfileBuilder()
}

// NewShareHoldersBuilder creates a new shareHolders builder instance
func NewShareHoldersBuilder() ShareHoldersBuilder {
	return createShareHoldersBuilder()
}

// NewShareHolderBuilder creates a new shareHolder builder instance
func NewShareHolderBuilder() ShareHolderBuilder {
	hashAdapter := hash.NewAdapter()
	return createShareHolderBuilder(hashAdapter)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithShareHolders(shareHolders ShareHolders) Builder
	WithConnections(connections Connections) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Seed() string
	ShareHolders() ShareHolders
	Connections() Connections
}

// ShareHoldersBuilder represents a shareholders builder
type ShareHoldersBuilder interface {
	Create() ShareHoldersBuilder
	WithShareHolders(shareHolders []ShareHolder) ShareHoldersBuilder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	All() []ShareHolder
	Fetch(gov governments.Government) (ShareHolder, error)
}

// ShareHolderBuilder represents a shareholder builder
type ShareHolderBuilder interface {
	Create() ShareHolderBuilder
	WithGovernment(gov governments.Government) ShareHolderBuilder
	WithPublic(public shareholders.ShareHolder) ShareHolderBuilder
	WithSigPK(sigPK signature.PrivateKey) ShareHolderBuilder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Government() governments.Government
	Public() shareholders.ShareHolder
	SigPK() signature.PrivateKey
}

// ConnectionsBuilder represents a connections builder
type ConnectionsBuilder interface {
	Create() ConnectionsBuilder
	WithConnections(connections []Connection) ConnectionsBuilder
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

// Repository represents an identity repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (Identity, error)
}

// Service represents a shareholders service
type Service interface {
	Insert(ins Identity, password string) error
	Update(origin Identity, updated Identity, password string) error
	UpdateWithPassword(origin Identity, updated Identity, originalPassword string, updatedPassword string) error
	Delete(ins Identity, password string) error
}
