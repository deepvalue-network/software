package containers

import (
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/databases"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/graphbases"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/sets"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables"
	"github.com/steve-care-software/products/libs/hash"
)

type transaction struct {
	graphbase graphbases.Transaction
	database  databases.Transaction
	table     tables.Transaction
	set       sets.Transaction
}

func createTransactionWithGraphbase(
	graphbase graphbases.Transaction,
) Transaction {
	return createTransactionInternally(graphbase, nil, nil, nil)
}

func createTransactionWithDatabase(
	database databases.Transaction,
) Transaction {
	return createTransactionInternally(nil, database, nil, nil)
}

func createTransactionWithTable(
	table tables.Transaction,
) Transaction {
	return createTransactionInternally(nil, nil, table, nil)
}

func createTransactionWithSet(
	set sets.Transaction,
) Transaction {
	return createTransactionInternally(nil, nil, nil, set)
}

func createTransactionInternally(
	graphbase graphbases.Transaction,
	database databases.Transaction,
	table tables.Transaction,
	set sets.Transaction,
) Transaction {
	out := transaction{
		graphbase: graphbase,
		database:  database,
		table:     table,
		set:       set,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	if obj.IsGraphbase() {
		return obj.Graphbase().Hash()
	}

	if obj.IsDatabase() {
		return obj.Database().Hash()
	}

	if obj.IsTable() {
		return obj.Table().Hash()
	}

	return obj.Set().Hash()
}

// IsGraphbase retruns true if there is a graphbase, false otherwise
func (obj *transaction) IsGraphbase() bool {
	return obj.graphbase != nil
}

// Graphbase returns the graphbase, if any
func (obj *transaction) Graphbase() graphbases.Transaction {
	return obj.graphbase
}

// IsDatabase retruns true if there is a database, false otherwise
func (obj *transaction) IsDatabase() bool {
	return obj.database != nil
}

// Database returns the database, if any
func (obj *transaction) Database() databases.Transaction {
	return obj.database
}

// IsDatabase retruns true if there is a database, false otherwise
func (obj *transaction) IsTable() bool {
	return obj.table != nil
}

// Table returns the database, if any
func (obj *transaction) Table() tables.Transaction {
	return obj.table
}

// IsSet retruns true if there is a set, false otherwise
func (obj *transaction) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *transaction) Set() sets.Transaction {
	return obj.set
}
