package deletes

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter          hash.Adapter
	graphbase            selectors.Selector
	mustBeGraphbaseEmpty bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:          hashAdapter,
		graphbase:            nil,
		mustBeGraphbaseEmpty: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithGraphbase adds a graphbase selector to the builder
func (app *builder) WithGraphbase(graphbase selectors.Selector) Builder {
	app.graphbase = graphbase
	return app
}

// MustBeGraphbaseEmpty flags the builder as must-be-database-empty
func (app *builder) MustBeGraphbaseEmpty() Builder {
	app.mustBeGraphbaseEmpty = true
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase selector is mandatory in order to build a Transaction instance")
	}

	mustBeGraphbaseEmpty := "false"
	if app.mustBeGraphbaseEmpty {
		mustBeGraphbaseEmpty = "true"
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.graphbase.Hash().Bytes(),
		[]byte(mustBeGraphbaseEmpty),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.graphbase, app.mustBeGraphbaseEmpty), nil
}
