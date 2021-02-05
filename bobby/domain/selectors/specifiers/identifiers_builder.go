package specifiers

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
)

type identifiersBuilder struct {
	hashAdapter hash.Adapter
	list        []Identifier
}

func createIdentifiersBuilder(hashAdapter hash.Adapter) IdentifiersBuilder {
	out := identifiersBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *identifiersBuilder) Create() IdentifiersBuilder {
	return createIdentifiersBuilder(app.hashAdapter)
}

// WithIdentifiers add identifiers to the builder
func (app *identifiersBuilder) WithIdentifiers(identifiers []Identifier) IdentifiersBuilder {
	app.list = identifiers
	return app
}

// Now builds a new Identifiers instance
func (app *identifiersBuilder) Now() (Identifiers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the identifiers are mandatory in order to build an Identifiers instance")
	}

	data := [][]byte{}
	for _, oneIdentifier := range app.list {
		data = append(data, oneIdentifier.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createIdentifiers(*hsh, app.list), nil
}
