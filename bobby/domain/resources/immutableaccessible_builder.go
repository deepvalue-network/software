package resources

import (
	"errors"

	"github.com/steve-care-software/products/libs/hash"
)

type immutableAccessibleBuilder struct {
	immutableBuilder ImmutableBuilder
	hash             *hash.Hash
	access           Access
}

func createImmutableAccessibleBuilder(
	immutableBuilder ImmutableBuilder,
) ImmutableAccessibleBuilder {
	out := immutableAccessibleBuilder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		access:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *immutableAccessibleBuilder) Create() ImmutableAccessibleBuilder {
	return createImmutableAccessibleBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *immutableAccessibleBuilder) WithHash(hash hash.Hash) ImmutableAccessibleBuilder {
	app.hash = &hash
	return app
}

// WithAccess adds an access to the builder
func (app *immutableAccessibleBuilder) WithAccess(access Access) ImmutableAccessibleBuilder {
	app.access = access
	return app
}

// Now builds a new ImmutableAccessible instance
func (app *immutableAccessibleBuilder) Now() (ImmutableAccessible, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an ImmutableAccessible instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).Now()
	if err != nil {
		return nil, err
	}

	if app.access != nil {
		return createImmutableAccessibleWithAccess(immutable, app.access), nil
	}

	return createImmutableAccessible(immutable), nil
}
