package deletes

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type transaction struct {
	hash           hash.Hash
	table          selectors.Selector
	mustBeRowEmpty bool
}

func createTransaction(
	hash hash.Hash,
	table selectors.Selector,
	mustBeRowEmpty bool,
) Transaction {
	out := transaction{
		hash:           hash,
		table:          table,
		mustBeRowEmpty: mustBeRowEmpty,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Table returns the table
func (obj *transaction) Table() selectors.Selector {
	return obj.table
}

// MustBeRowEmpty returns true if it must be row empty, false otherwise
func (obj *transaction) MustBeRowEmpty() bool {
	return obj.mustBeRowEmpty
}
