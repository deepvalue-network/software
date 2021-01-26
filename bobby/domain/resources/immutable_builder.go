package resources

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/hash"
)

type immutableBuilder struct {
	id        *uuid.UUID
	hash      *hash.Hash
	createdOn *time.Time
}

func createImmutableBuilder() ImmutableBuilder {
	out := immutableBuilder{
		id:        nil,
		hash:      nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *immutableBuilder) Create() ImmutableBuilder {
	return createImmutableBuilder()
}

// WithID adds an ID to the builder
func (app *immutableBuilder) WithID(id *uuid.UUID) ImmutableBuilder {
	app.id = id
	return app
}

// WithHash adds an hash to the builder
func (app *immutableBuilder) WithHash(hash hash.Hash) ImmutableBuilder {
	app.hash = &hash
	return app
}

// Now builds a new Immutable instance
func (app *immutableBuilder) Now() (Immutable, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Immutable instance")
	}

	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	return createImmutable(&id, *app.hash, createdOn), nil
}
