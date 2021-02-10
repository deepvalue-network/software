package peers

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// NormalProtocol represents the normal protocol
const NormalProtocol = "https"

// TorProtocol represents the tor protocol
const TorProtocol = "tor"

const protocolSeparator = "://"
const protocolPattern = "%s%s%s"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPeerBuilder creates a new peer builder instance
func NewPeerBuilder() PeerBuilder {
	return createPeerBuilder()
}

// NewPointer returns a new peers pointer
func NewPointer() *peers {
	return new(peers)
}

// NewPeerPointer returns a new peer pointer
func NewPeerPointer() *peer {
	return new(peer)
}

// Builder represents a peers builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithSyncDuration(syncDuration time.Duration) Builder
	WithList(list []Peer) Builder
	LastSyncTime(lastSyncTime time.Time) Builder
	Now() (Peers, error)
}

// Peers represents a peer list
type Peers interface {
	ID() *uuid.UUID
	SyncInterval() time.Duration
	All() []Peer
	Add(ins Peer) error
	Merge(ins Peers)
	Delete(ins Peer) error
	HasLastSync() bool
	LastSync() *time.Time
}

// PeerBuilder represents a peers builder
type PeerBuilder interface {
	Create() PeerBuilder
	WithOriginal(original Peer) PeerBuilder
	WithServer(server string) PeerBuilder
	CreatedOn(createdOn time.Time) PeerBuilder
	LastUpdatedOn(lastUpdatedOn time.Time) PeerBuilder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	Content() Content
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
}

// Content represents a peer content
type Content interface {
	String() string
	IsNormal() bool
	Normal() Server
	IsTor() bool
	Tor() Server
}

// Server represents a peer server
type Server interface {
	Host() string
	Port() uint
	String() string
}
