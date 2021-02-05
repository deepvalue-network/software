package resources

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	immutableBuilder ImmutableBuilder
	mutableBuilder   MutableBuilder
	hash             *hash.Hash
	accessible       Accessible
	immutable        Immutable
	mutable          Mutable
}

func createBuilder(
	immutableBuilder ImmutableBuilder,
	mutableBuilder MutableBuilder,
) Builder {
	out := builder{
		immutableBuilder: immutableBuilder,
		mutableBuilder:   mutableBuilder,
		hash:             nil,
		accessible:       nil,
		immutable:        nil,
		mutable:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.immutableBuilder,
		app.mutableBuilder,
	)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithAccessible adds an accessible to the builder
func (app *builder) WithAccessible(accessible Accessible) Builder {
	app.accessible = accessible
	return app
}

// WithImmutable adds an immutable to the builder
func (app *builder) WithImmutable(immutable Immutable) Builder {
	app.immutable = immutable
	return app
}

// WithMutable adds a mutable to the builder
func (app *builder) WithMutable(mutable Mutable) Builder {
	app.mutable = mutable
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.hash != nil && app.accessible != nil {
		if app.accessible.IsMutable() {
			mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).Now()
			if err != nil {
				return nil, err
			}

			app.mutable = mutable
		}

		if app.accessible.IsImmutable() {
			immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).Now()
			if err != nil {
				return nil, err
			}

			app.immutable = immutable
		}
	}

	if app.immutable != nil {
		return createResourceWithImmutable(app.immutable), nil
	}

	if app.mutable != nil {
		return createResourceWithMutable(app.mutable), nil
	}

	return nil, errors.New("the Resource is invalid")
}
