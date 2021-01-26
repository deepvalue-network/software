package entities

import (
	"errors"
	"time"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

type immutableBuilder struct {
	hash      *hash.Hash
	createdOn *time.Time
}

func createImmutableBuilder() ImmutableBuilder {
	out := immutableBuilder{
		hash:      nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *immutableBuilder) Create() ImmutableBuilder {
	return createImmutableBuilder()
}

// WithHash adds an hash to the builder
func (app *immutableBuilder) WithHash(hash hash.Hash) ImmutableBuilder {
	app.hash = &hash
	return app
}

// CreatedOn adds a creation time to the builder
func (app *immutableBuilder) CreatedOn(createdOn *time.Time) ImmutableBuilder {
	app.createdOn = createdOn
	return app
}

// Now builds a new Immutable instance
func (app *immutableBuilder) Now() (Immutable, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Immutable instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	return createImmutable(*app.hash, *app.createdOn), nil
}
