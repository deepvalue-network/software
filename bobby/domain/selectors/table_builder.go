package selectors

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type tableBuilder struct {
	hashAdapter hash.Adapter
	graphbase   specifiers.Specifier
	db          specifiers.Specifier
	schema      schemas.Schema
	specifier   specifiers.Specifier
}

func createTableBuilder(
	hashAdapter hash.Adapter,
) TableBuilder {
	out := tableBuilder{
		hashAdapter: hashAdapter,
		graphbase:   nil,
		db:          nil,
		schema:      nil,
		specifier:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *tableBuilder) Create() TableBuilder {
	return createTableBuilder(app.hashAdapter)
}

// WithGraphbase adds a graphbase to the builder
func (app *tableBuilder) WithGraphbase(graphbase specifiers.Specifier) TableBuilder {
	app.graphbase = graphbase
	return app
}

// WithDatabase adds a database to the builder
func (app *tableBuilder) WithDatabase(db specifiers.Specifier) TableBuilder {
	app.db = db
	return app
}

// WithSchema adds a schema to the builder
func (app *tableBuilder) WithSchema(schema schemas.Schema) TableBuilder {
	app.schema = schema
	return app
}

// WithSpecifier adds a specifier to the builder
func (app *tableBuilder) WithSpecifier(specifier specifiers.Specifier) TableBuilder {
	app.specifier = specifier
	return app
}

// Now builds a new Table instance
func (app *tableBuilder) Now() (Table, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Table instance")
	}

	if app.db == nil {
		return nil, errors.New("the database is mandatory in order to build a Table instance")
	}

	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Table instance")
	}

	var content TableContent
	if app.specifier != nil {
		content = createTableContentWithSpecifier(app.specifier)
	}

	if content == nil {
		return nil, errors.New("the content (specifier) is mandatory in order to build a Table instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.graphbase.Hash().Bytes(),
		app.db.Hash().Bytes(),
		app.schema.Resource().Hash().Bytes(),
		content.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createTable(*hsh, app.graphbase, app.db, app.schema, content), nil
}
