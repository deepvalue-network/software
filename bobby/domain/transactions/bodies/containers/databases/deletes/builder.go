package deletes

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	db               selectors.Selector
	mustBeTableEmpty bool
	mustBeSetEmpty   bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		db:               nil,
		mustBeTableEmpty: false,
		mustBeSetEmpty:   false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithDatabase adds the database selector to the builder
func (app *builder) WithDatabase(db selectors.Selector) Builder {
	app.db = db
	return app
}

// MustBeTableEmpty flags the builder as must-be-table-empty
func (app *builder) MustBeTableEmpty() Builder {
	app.mustBeTableEmpty = true
	return app
}

// MustBeSetEmpty flags the builder as must-be-set-empty
func (app *builder) MustBeSetEmpty() Builder {
	app.mustBeSetEmpty = true
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.db == nil {
		return nil, errors.New("the database selector is mandatory in order to build a Transaction instance")
	}

	mustBeTableEmpty := "false"
	if app.mustBeTableEmpty {
		mustBeTableEmpty = "true"
	}

	mustBeSetEmpty := "false"
	if app.mustBeSetEmpty {
		mustBeSetEmpty = "true"
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.db.Hash().Bytes(),
		[]byte(mustBeTableEmpty),
		[]byte(mustBeSetEmpty),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.db, app.mustBeTableEmpty, app.mustBeSetEmpty), nil
}
