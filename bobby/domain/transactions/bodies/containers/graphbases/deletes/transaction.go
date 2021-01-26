package deletes

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type transaction struct {
	hash                 hash.Hash
	graphbase            selectors.Selector
	mustBeGraphbaseEmpty bool
}

func createTransaction(
	hash hash.Hash,
	graphbase selectors.Selector,
	mustBeGraphbaseEmpty bool,
) Transaction {
	out := transaction{
		hash:                 hash,
		graphbase:            graphbase,
		mustBeGraphbaseEmpty: mustBeGraphbaseEmpty,
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

// MustBeGraphbaseEmpty returns true if it must be database empty, false otherwise
func (obj *transaction) MustBeGraphbaseEmpty() bool {
	return obj.mustBeGraphbaseEmpty
}
