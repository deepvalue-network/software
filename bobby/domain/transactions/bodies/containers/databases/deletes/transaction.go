package deletes

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type transaction struct {
	hash             hash.Hash
	db               selectors.Selector
	mustBeTableEmpty bool
	mustBeSetEmpty   bool
}

func createTransaction(
	hash hash.Hash,
	db selectors.Selector,
	mustBeTableEmpty bool,
	mustBeSetEmpty bool,
) Transaction {
	out := transaction{
		hash:             hash,
		db:               db,
		mustBeTableEmpty: mustBeTableEmpty,
		mustBeSetEmpty:   mustBeSetEmpty,
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

// MustBeTableEmpty returns true if the database must be table-empty, false otherwise
func (obj *transaction) MustBeTableEmpty() bool {
	return obj.mustBeTableEmpty
}

// MustBeSetEmpty returns true if the database must be set-empty, false otherwise
func (obj *transaction) MustBeSetEmpty() bool {
	return obj.mustBeSetEmpty
}
