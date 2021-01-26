package peers

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type builder struct {
	syncInterval *time.Duration
}

func createBuilder() Builder {
	out := builder{
		syncInterval: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSyncDuration adds a sync duration to the builder
func (app *builder) WithSyncDuration(syncDuration time.Duration) Builder {
	app.syncInterval = &syncDuration
	return app
}

// Now builds a new Peers instance
func (app *builder) Now() (Peers, error) {
	if app.syncInterval == nil {
		return nil, errors.New("the sync interval is mandatory in order to build a Peers instance")
	}

	id := uuid.NewV4()
	return createPeers(&id, *app.syncInterval), nil
}
