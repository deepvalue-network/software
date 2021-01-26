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

// Builder represents a peers builder
type Builder interface {
	Create() Builder
	WithSyncDuration(syncDuration time.Duration) Builder
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
	WithNormalServer(normal string) PeerBuilder
	WithTorServer(tor string) PeerBuilder
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
