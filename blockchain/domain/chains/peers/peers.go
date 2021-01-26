package peers

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type peers struct {
	id           *uuid.UUID
	syncInterval time.Duration
	mp           map[string]Peer
	lst          []Peer
	lastSync     *time.Time
}

func createPeers(
	id *uuid.UUID,
	syncInterval time.Duration,
) Peers {
	out := peers{
		id:           id,
		syncInterval: syncInterval,
		mp:           map[string]Peer{},
		lst:          []Peer{},
		lastSync:     nil,
	}

	return &out
}

// ID returns the id
func (obj *peers) ID() *uuid.UUID {
	return obj.id
}

// SyncInterval returns the sync interval
func (obj *peers) SyncInterval() time.Duration {
	return obj.syncInterval
}

// All returns the peers
func (obj *peers) All() []Peer {
	return obj.lst
}

// Add adds a peer
func (obj *peers) Add(ins Peer) error {
	return nil
}

// Merge merges peers
func (obj *peers) Merge(ins Peers) {

}

// Delete deletes a peer
func (obj *peers) Delete(ins Peer) error {
	return nil
}

// HasLastSync returns true if there is a lastSync time, false otherwise
func (obj *peers) HasLastSync() bool {
	return obj.lastSync != nil
}

// LastSync returns the lastSync, if any
func (obj *peers) LastSync() *time.Time {
	return obj.lastSync
}
