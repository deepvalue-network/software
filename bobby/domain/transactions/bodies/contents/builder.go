package contents

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/rows"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	table       selectors.Selector
	rows        rows.Rows
	set         selectors.Selector
	elements    sets.Elements
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		table:       nil,
		rows:        nil,
		set:         nil,
		elements:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table selectors.Selector) Builder {
	app.table = table
	return app
}

// WithTableRows adds table rows to the builder
func (app *builder) WithTableRows(rows rows.Rows) Builder {
	app.rows = rows
	return app
}

// WithSet adds a set to the builder
func (app *builder) WithSet(set selectors.Selector) Builder {
	app.set = set
	return app
}

// WithSetElements adds set elements to the builder
func (app *builder) WithSetElements(elements sets.Elements) Builder {
	app.elements = elements
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {

	data := [][]byte{}
	if app.table != nil {
		data = append(data, app.table.Hash().Bytes())
	}

	if app.rows != nil {
		data = append(data, app.rows.Hash().Bytes())
	}

	if app.set != nil {
		data = append(data, app.set.Hash().Bytes())
	}

	if app.elements != nil {
		data = append(data, app.elements.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.table != nil && app.rows != nil {
		table := createTable(app.table, app.rows)
		return createTransactionWithTable(*hsh, table), nil
	}

	if app.set != nil && app.elements != nil {
		set := createSet(app.set, app.elements)
		return createTransactionWithSet(*hsh, set), nil
	}

	return nil, errors.New("the Transaction is invalid")
}
