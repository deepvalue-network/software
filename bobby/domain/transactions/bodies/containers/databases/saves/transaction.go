package saves

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type transaction struct {
	hash      hash.Hash
	graphbase selectors.Selector
	name      string
	database  selectors.Selector
}

func createTransaction(
	hash hash.Hash,
	graphbase selectors.Selector,
	name string,
) Transaction {
	return createTransactionInternally(hash, graphbase, name, nil)
}

func createTransactionWithDatabase(
	hash hash.Hash,
	graphbase selectors.Selector,
	name string,
	database selectors.Selector,
) Transaction {
	return createTransactionInternally(hash, graphbase, name, database)
}

func createTransactionInternally(
	hash hash.Hash,
	graphbase selectors.Selector,
	name string,
	database selectors.Selector,
) Transaction {
	out := transaction{
		hash:      hash,
		graphbase: graphbase,
		name:      name,
		database:  database,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Graphbase returns the graphbase
func (obj *transaction) Graphbase() selectors.Selector {
	return obj.graphbase
}

// Name returns the name
func (obj *transaction) Name() string {
	return obj.name
}

// HasDatabase returns true if there is a database, false otehrwise
func (obj *transaction) HasDatabase() bool {
	return obj.database != nil
}

// Database returns the database, if any
func (obj *transaction) Database() selectors.Selector {
	return obj.database
}
