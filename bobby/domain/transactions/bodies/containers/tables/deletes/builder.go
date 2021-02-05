package deletes

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	table          selectors.Selector
	mustBeRowEmpty bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		table:          nil,
		mustBeRowEmpty: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTable adds a table selector to the builder
func (app *builder) WithTable(table selectors.Selector) Builder {
	app.table = table
	return app
}

// MustBeRowEmpty flags the builder as must-be-row-empty
func (app *builder) MustBeRowEmpty() Builder {
	app.mustBeRowEmpty = true
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.table == nil {
		return nil, errors.New("the table selector is mandatory in order to build a Transaction instance")
	}

	mustBeEmptyRow := "false"
	if app.mustBeRowEmpty {
		mustBeEmptyRow = "true"
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.table.Hash().Bytes(),
		[]byte(mustBeEmptyRow),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.table, app.mustBeRowEmpty), nil
}
