package deletes

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	set                selectors.Selector
	mustBeElementEmpty bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		set:                nil,
		mustBeElementEmpty: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithSet adds a set selector to the builder
func (app *builder) WithSet(set selectors.Selector) Builder {
	app.set = set
	return app
}

// MustBeElementEmpty flags the builder as must-be-element-empty
func (app *builder) MustBeElementEmpty() Builder {
	app.mustBeElementEmpty = true
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.set == nil {
		return nil, errors.New("the set selector is mandatory in order to build a Transaction instance")
	}

	mustBeElementEmpty := "false"
	if app.mustBeElementEmpty {
		mustBeElementEmpty = "true"
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.set.Hash().Bytes(),
		[]byte(mustBeElementEmpty),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.set, app.mustBeElementEmpty), nil
}
