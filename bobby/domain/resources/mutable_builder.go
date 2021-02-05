package resources

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/hash"
)

type mutableBuilder struct {
	immutableBuilder ImmutableBuilder
	hash             *hash.Hash
	parent           Mutable
}

func createMutableBuilder(
	immutableBuilder ImmutableBuilder,
) MutableBuilder {
	out := mutableBuilder{
		immutableBuilder: immutableBuilder,
		hash:             nil,
		parent:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *mutableBuilder) Create() MutableBuilder {
	return createMutableBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *mutableBuilder) WithHash(hash hash.Hash) MutableBuilder {
	app.hash = &hash
	return app
}

// WithParent adds a parent to the builder
func (app *mutableBuilder) WithParent(parent Mutable) MutableBuilder {
	app.parent = parent
	return app
}

// Now builds a new Mutable instance
func (app *mutableBuilder) Now() (Mutable, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Mutable instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).Now()
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		if !app.parent.CreatedOn().Before(immutable.CreatedOn()) {
			str := fmt.Sprintf("the parent's creation time (%s) must be before the creation time (%s)", app.parent.CreatedOn().String(), immutable.CreatedOn().String())
			return nil, errors.New(str)
		}

		return createMutableWithParent(immutable, app.parent), nil
	}

	return createMutable(immutable), nil
}
