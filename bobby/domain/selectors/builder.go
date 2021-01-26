package selectors

import (
	"errors"

	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	pkAdapter     encryption.Adapter
	decryptionKey encryption.PrivateKey
	graphbase     Graphbase
	db            Database
	table         Table
	set           Set
}

func createBuilder(
	hashAdapter hash.Adapter,
	pkAdapter encryption.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		pkAdapter:     pkAdapter,
		decryptionKey: nil,
		graphbase:     nil,
		db:            nil,
		table:         nil,
		set:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.pkAdapter,
	)
}

// WithDecryptionKey adds a decryptionKey to the builder
func (app *builder) WithDecryptionKey(decryptionKey encryption.PrivateKey) Builder {
	app.decryptionKey = decryptionKey
	return app
}

// WithGraphbase adds a graphbase to the builder
func (app *builder) WithGraphbase(graphbase Graphbase) Builder {
	app.graphbase = graphbase
	return app
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(db Database) Builder {
	app.db = db
	return app
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table Table) Builder {
	app.table = table
	return app
}

// WithSet adds a set to the builder
func (app *builder) WithSet(set Set) Builder {
	app.set = set
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.decryptionKey != nil {
		return nil, errors.New("the decryptionKey is mandatory in order to build a Selector instance")
	}

	data := [][]byte{
		app.pkAdapter.ToBytes(app.decryptionKey),
	}

	var content Content
	if app.graphbase != nil {
		data = append(data, app.graphbase.Hash().Bytes())
		content = createContentWithGraphbase(app.graphbase)
	}

	if app.db != nil {
		data = append(data, app.db.Hash().Bytes())
		content = createContentWithDatabase(app.db)
	}

	if app.table != nil {
		data = append(data, app.table.Hash().Bytes())
		content = createContentWithTable(app.table)
	}

	if app.set != nil {
		data = append(data, app.set.Hash().Bytes())
		content = createContentWithSet(app.set)
	}

	if content == nil {
		return nil, errors.New("the content (graphbase, database, table, set) is mandatory in order to build a Selector instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createSelector(*hsh, app.decryptionKey, content), nil
}
