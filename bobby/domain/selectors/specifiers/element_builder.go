package specifiers

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	hash        *hash.Hash
	id          *uuid.UUID
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		hash:        nil,
		id:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(app.hashAdapter)
}

// WithHash adds an hash to the builder
func (app *elementBuilder) WithHash(hash hash.Hash) ElementBuilder {
	app.hash = &hash
	return app
}

// WithID adds an id to the builder
func (app *elementBuilder) WithID(id *uuid.UUID) ElementBuilder {
	app.id = id
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.hash != nil {
		hsh, err := app.hashAdapter.FromBytes(app.hash.Bytes())
		if err != nil {
			return nil, err
		}

		return createElementWithHash(*hsh, app.hash), nil
	}

	if app.id != nil {
		hsh, err := app.hashAdapter.FromBytes(app.id.Bytes())
		if err != nil {
			return nil, err
		}

		return createElementWithID(*hsh, app.id), nil
	}

	return nil, errors.New("the Element is invalid")
}
