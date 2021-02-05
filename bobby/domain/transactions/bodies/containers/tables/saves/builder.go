package saves

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	db          selectors.Selector
	schema      schemas.Schema
	table       selectors.Selector
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		db:          nil,
		schema:      nil,
		table:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(db selectors.Selector) Builder {
	app.db = db
	return app
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema schemas.Schema) Builder {
	app.schema = schema
	return app
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table selectors.Selector) Builder {
	app.table = table
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.db == nil {
		return nil, errors.New("the database is mandatory in order to build a Transaction instance")
	}

	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Transaction instance")
	}

	data := [][]byte{
		app.db.Hash().Bytes(),
		app.schema.Resource().Hash().Bytes(),
	}

	if app.table != nil {
		data = append(data, app.table.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.table != nil {
		return createTransactionWithTable(*hsh, app.db, app.schema, app.table), nil
	}

	return createTransaction(*hsh, app.db, app.schema), nil
}
