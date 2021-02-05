package containers

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/databases"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/graphbases"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/sets"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers/tables"
)

type builder struct {
	graphbase graphbases.Transaction
	database  databases.Transaction
	table     tables.Transaction
	set       sets.Transaction
}

func createBuilder() Builder {
	out := builder{
		graphbase: nil,
		database:  nil,
		table:     nil,
		set:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGraphbase adds a graphbase to the builder
func (app *builder) WithGraphbase(graph graphbases.Transaction) Builder {
	app.graphbase = graph
	return app
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(db databases.Transaction) Builder {
	app.database = db
	return app
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table tables.Transaction) Builder {
	app.table = table
	return app
}

// WithSet adds a set to the builder
func (app *builder) WithSet(set sets.Transaction) Builder {
	app.set = set
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.graphbase != nil {
		return createTransactionWithGraphbase(app.graphbase), nil
	}

	if app.database != nil {
		return createTransactionWithDatabase(app.database), nil
	}

	if app.table != nil {
		return createTransactionWithTable(app.table), nil
	}

	if app.set != nil {
		return createTransactionWithSet(app.set), nil
	}

	return nil, errors.New("the Transaction is invalid")
}
