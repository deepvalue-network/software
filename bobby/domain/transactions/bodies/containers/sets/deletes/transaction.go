package deletes

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash               hash.Hash
	set                selectors.Selector
	mustBeElementEmpty bool
}

func createTransaction(
	hash hash.Hash,
	set selectors.Selector,
	mustBeElementEmpty bool,
) Transaction {
	out := transaction{
		hash:               hash,
		set:                set,
		mustBeElementEmpty: mustBeElementEmpty,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Set returns the set
func (obj *transaction) Set() selectors.Selector {
	return obj.set
}

// MustBeElementEmpty returns true if it must be element empty, false otherwise
func (obj *transaction) MustBeElementEmpty() bool {
	return obj.mustBeElementEmpty
}
