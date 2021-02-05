package owners

import (
	"errors"

	"github.com/deepvalue-network/software/diamonds/domain/genesis/spends/views"
	"github.com/deepvalue-network/software/diamonds/domain/owners"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	owner       owners.Owner
	genesis     views.Genesis
}

func createBuilder(hashAdapter hash.Adapter) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		owner:       nil,
		genesis:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithOwner adds an owner to the builder
func (app *builder) WithOwner(owner owners.Owner) Builder {
	app.owner = owner
	return app
}

// WithGenesis adds a genesis to the builder
func (app *builder) WithGenesis(genesis views.Genesis) Builder {
	app.genesis = genesis
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.genesis == nil {
		return nil, errors.New("the view genesis is mandatory in order to build an owner Genesis instance")
	}

	if app.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a Genesis instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.genesis.Hash().Bytes(),
		app.owner.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(*hsh, app.owner, app.genesis), nil
}
