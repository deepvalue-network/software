package resources

import (
	"errors"

	"github.com/steve-care-software/products/libs/hash"
)

type mutableAccessibleBuilder struct {
	mutableBuilder MutableBuilder
	hash           *hash.Hash
	parent         Mutable
	access         Access
}

func createMutableAccessibleBuilder(
	mutableBuilder MutableBuilder,
) MutableAccessibleBuilder {
	out := mutableAccessibleBuilder{
		mutableBuilder: mutableBuilder,
		hash:           nil,
		parent:         nil,
		access:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *mutableAccessibleBuilder) Create() MutableAccessibleBuilder {
	return createMutableAccessibleBuilder(app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *mutableAccessibleBuilder) WithHash(hash hash.Hash) MutableAccessibleBuilder {
	app.hash = &hash
	return app
}

// WithParent adds a parent to the builder
func (app *mutableAccessibleBuilder) WithParent(parent Mutable) MutableAccessibleBuilder {
	app.parent = parent
	return app
}

// WithAccess adds an access to the builder
func (app *mutableAccessibleBuilder) WithAccess(access Access) MutableAccessibleBuilder {
	app.access = access
	return app
}

// Now builds a new MutableAccessible instance
func (app *mutableAccessibleBuilder) Now() (MutableAccessible, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build an MutableAccessible instance")
	}

	mutableBuilder := app.mutableBuilder.Create().WithHash(*app.hash)
	if app.parent != nil {
		mutableBuilder.WithParent(app.parent)
	}

	mutable, err := mutableBuilder.Now()
	if err != nil {
		return nil, err
	}

	if app.access != nil {
		return createMutableAccessibleWithAccess(mutable, app.access), nil
	}

	return createMutableAccessible(mutable), nil
}
