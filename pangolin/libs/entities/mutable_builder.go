package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

type mutableBuilder struct {
	hash          *hash.Hash
	createdOn     *time.Time
	lastUpdatedOn *time.Time
}

func createMutableBuilder() MutableBuilder {
	out := mutableBuilder{
		hash:          nil,
		createdOn:     nil,
		lastUpdatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *mutableBuilder) Create() MutableBuilder {
	return createMutableBuilder()
}

// WithHash adds an hash to the builder
func (app *mutableBuilder) WithHash(hash hash.Hash) MutableBuilder {
	app.hash = &hash
	return app
}

// CreatedOn adds a creation time
func (app *mutableBuilder) CreatedOn(createdOn *time.Time) MutableBuilder {
	app.createdOn = createdOn
	return app
}

// LastUpdatedOn adds a lastUpdatedOn time
func (app *mutableBuilder) LastUpdatedOn(lastUpdatedOn *time.Time) MutableBuilder {
	app.lastUpdatedOn = lastUpdatedOn
	return app
}

// Now builds a new Mutable instance
func (app *mutableBuilder) Now() (Mutable, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Mutable instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.lastUpdatedOn == nil {
		lastUpdatedOn := time.Now().UTC()
		app.lastUpdatedOn = &lastUpdatedOn
	}

	if app.createdOn.After(*app.lastUpdatedOn) {
		str := fmt.Sprintf("the creation time (%s) cannot be after the lastUpdated time (%s)", app.createdOn.String(), app.lastUpdatedOn.String())
		return nil, errors.New(str)
	}

	return createMutable(*app.hash, *app.createdOn, *app.lastUpdatedOn), nil
}
