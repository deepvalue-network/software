package saves

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash   hash.Hash
	db     selectors.Selector
	schema schemas.Schema
	set    selectors.Selector
}

func createTransaction(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
) Transaction {
	return createTransactionInternally(hash, db, schema, nil)
}

func createTransactionWithSet(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
	set selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, db, schema, set)
}

func createTransactionInternally(
	hash hash.Hash,
	db selectors.Selector,
	schema schemas.Schema,
	set selectors.Selector,
) Transaction {
	out := transaction{
		hash:   hash,
		db:     db,
		schema: schema,
		set:    set,
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

// HasSet returns true if there is a set, false otherwise
func (obj *transaction) HasSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *transaction) Set() selectors.Selector {
	return obj.set
}
