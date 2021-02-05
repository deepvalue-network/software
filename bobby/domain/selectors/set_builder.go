package selectors

import (
	"errors"
	"strconv"

	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type setBuilder struct {
	hashAdapter hash.Adapter
	graphbase   specifiers.Specifier
	db          specifiers.Specifier
	schema      schemas.Schema
	specifier   specifiers.Specifier
	from        *uint
	to          *uint
}

func createSetBuilder(
	hashAdapter hash.Adapter,
) SetBuilder {
	out := setBuilder{
		hashAdapter: hashAdapter,
		graphbase:   nil,
		db:          nil,
		schema:      nil,
		specifier:   nil,
		from:        nil,
		to:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *setBuilder) Create() SetBuilder {
	return createSetBuilder(app.hashAdapter)
}

// WithGraphbase adds a graphbase to the builder
func (app *setBuilder) WithGraphbase(graphbase specifiers.Specifier) SetBuilder {
	app.graphbase = graphbase
	return app
}

// WithDatabase adds a database to the builder
func (app *setBuilder) WithDatabase(db specifiers.Specifier) SetBuilder {
	app.db = db
	return app
}

// WithSchema adds a schema to the builder
func (app *setBuilder) WithSchema(schema schemas.Schema) SetBuilder {
	app.schema = schema
	return app
}

// WithSpecifier adds a specifier to the builder
func (app *setBuilder) WithSpecifier(specifier specifiers.Specifier) SetBuilder {
	app.specifier = specifier
	return app
}

// From adds a from to the builder
func (app *setBuilder) From(from uint) SetBuilder {
	app.from = &from
	return app
}

// To adds a to to the builder
func (app *setBuilder) To(to uint) SetBuilder {
	app.to = &to
	return app
}

// Now builds a new Set instance
func (app *setBuilder) Now() (Set, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Set instance")
	}

	if app.db == nil {
		return nil, errors.New("the database is mandatory in order to build a Set instance")
	}

	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Set instance")
	}

	data := [][]byte{
		app.graphbase.Hash().Bytes(),
		app.db.Hash().Bytes(),
		app.schema.Resource().Hash().Bytes(),
	}

	var content SetContent
	if app.specifier != nil {
		data = append(data, app.specifier.Hash().Bytes())
		content = createSetContentWithSpecifier(app.specifier)
	}

	if app.from != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.from))))
	}

	if app.to != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.to))))
	}

	if content == nil {
		return nil, errors.New("the content (specifier) is mandatory in order to build a Set instance")
	}

	var rank SetRank
	if app.from != nil && app.to != nil {
		rank = createSetRankWithFromAndTo(app.from, app.to)
	}

	if rank != nil && app.from != nil {
		rank = createSetRankWithFrom(app.from)
	}

	if rank != nil && app.to != nil {
		rank = createSetRankWithTo(app.to)
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if rank != nil {
		return createSetWithRank(*hsh, app.graphbase, app.db, app.schema, content, rank), nil
	}

	return createSet(*hsh, app.graphbase, app.db, app.schema, content), nil
}
