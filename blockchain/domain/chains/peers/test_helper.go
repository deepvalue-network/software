package peers

import (
	"fmt"
	"time"
)

// CreatePeerForTests creates a new peer instance for tests
func CreatePeerForTests() Peer {
	normal := fmt.Sprintf("%s://127.0.0.1:80", NormalProtocol)
	ins, err := NewPeerBuilder().Create().WithServer(normal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreatePeersForTests creates a new peers instance for tests
func CreatePeersForTests() Peers {
	first := CreatePeerForTests()
	list := []Peer{
		first,
	}

	lastSyncTime := time.Now().UTC()
	syncDuration := time.Duration(time.Second * 5)
	ins, err := NewBuilder().Create().WithList(list).WithSyncDuration(syncDuration).LastSyncTime(lastSyncTime).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
