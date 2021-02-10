package peers

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type peers struct {
	id           *uuid.UUID    `hydro:"ID, ID"`
	syncInterval time.Duration `hydro:"SyncInterval, SyncInterval"`
	mp           map[string]Peer
	lst          []Peer     `hydro:"All, List"`
	lastSync     *time.Time `hydro:"LastSync, LastSyncTime"`
}

func createPeers(
	id *uuid.UUID,
	syncInterval time.Duration,
	mp map[string]Peer,
	lst []Peer,
) Peers {
	return createPeersInternally(id, syncInterval, mp, lst, nil)
}

func createPeersWithLastSync(
	id *uuid.UUID,
	syncInterval time.Duration,
	mp map[string]Peer,
	lst []Peer,
	lastSync *time.Time,
) Peers {
	return createPeersInternally(id, syncInterval, mp, lst, lastSync)
}

func createPeersInternally(
	id *uuid.UUID,
	syncInterval time.Duration,
	mp map[string]Peer,
	lst []Peer,
	lastSync *time.Time,
) Peers {
	out := peers{
		id:           id,
		syncInterval: syncInterval,
		mp:           mp,
		lst:          lst,
		lastSync:     lastSync,
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
	keyname := ins.Content().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the peer (host: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.mp[keyname] = ins
	obj.lst = append(obj.lst, ins)
	return nil
}

// Merge merges peers
func (obj *peers) Merge(ins Peers) {
	all := ins.All()
	for _, onePeer := range all {
		obj.Add(onePeer)
	}
}

// Delete deletes a peer
func (obj *peers) Delete(ins Peer) error {
	keyname := ins.Content().String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the peer (host: %s) does not exists", keyname)
		return errors.New(str)
	}

	delete(obj.mp, keyname)
	for index, onePeer := range obj.lst {
		peerKey := onePeer.Content().String()
		if peerKey != keyname {
			continue
		}

		obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
	}

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
