package saves

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashadapter hash.Adapter
	db          selectors.Selector
	schema      schemas.Schema
	set         selectors.Selector
}

func createBuilder(
	hashadapter hash.Adapter,
) Builder {
	out := builder{
		hashadapter: hashadapter,
		db:          nil,
		schema:      nil,
		set:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashadapter)
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

// WithSet adds a set to the builder
func (app *builder) WithSet(set selectors.Selector) Builder {
	app.set = set
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

	if app.set != nil {
		data = append(data, app.set.Hash().Bytes())
	}

	hsh, err := app.hashadapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.set != nil {
		return createTransactionWithSet(*hsh, app.db, app.schema, app.set), nil
	}

	return createTransaction(*hsh, app.db, app.schema), nil
}
