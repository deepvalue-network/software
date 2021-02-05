package saves

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash   hash.Hash
	db     selectors.Selector
	schema schemas.Schema
	table  selectors.Selector
}

func createTransaction(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
) Transaction {
	return createTransactionInternally(hash, db, schema, nil)
}

func createTransactionWithTable(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
	table selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, db, schema, table)
}

func createTransactionInternally(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
	table selectors.Selector,
) Transaction {
	out := transaction{
		hash:   hash,
		db:     db,
		schema: schema,
		table:  table,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Database returns the database selector
func (obj *transaction) Database() selectors.Selector {
	return obj.db
}

// Schema returns the schema
func (obj *transaction) Schema() schemas.Schema {
	return obj.schema
}

// HasTable returns true if there is a table, false otherwise
func (obj *transaction) HasTable() bool {
	return obj.table != nil
}

// Table returns the table, if any
func (obj *transaction) Table() selectors.Selector {
	return obj.table
}
