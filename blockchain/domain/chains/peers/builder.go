package peers

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type builder struct {
	id           *uuid.UUID
	syncInterval *time.Duration
	list         []Peer
	lastSyncTime *time.Time
}

func createBuilder() Builder {
	out := builder{
		id:           nil,
		syncInterval: nil,
		list:         nil,
		lastSyncTime: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithSyncDuration adds a sync duration to the builder
func (app *builder) WithSyncDuration(syncDuration time.Duration) Builder {
	app.syncInterval = &syncDuration
	return app
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Peer) Builder {
	app.list = list
	return app
}

// LastSyncTime adds a last sync time to the builder
func (app *builder) LastSyncTime(lastSyncTime time.Time) Builder {
	app.lastSyncTime = &lastSyncTime
	return app
}

// Now builds a new Peers instance
func (app *builder) Now() (Peers, error) {
	if app.id == nil {
		id := uuid.NewV4()
		app.id = &id
	}

	if app.syncInterval == nil {
		return nil, errors.New("the sync interval is mandatory in order to build a Peers instance")
	}

	if app.list == nil {
		app.list = []Peer{}
	}

	mp := map[string]Peer{}
	for _, onePeer := range app.list {
		keyname := onePeer.Content().String()
		mp[keyname] = onePeer
	}

	if app.lastSyncTime != nil {
		return createPeersWithLastSync(app.id, *app.syncInterval, mp, app.list, app.lastSyncTime), nil
	}

	return createPeers(app.id, *app.syncInterval, mp, app.list), nil
}
