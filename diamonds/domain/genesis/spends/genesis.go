package spends

import (
	"time"

	domain_genesis "github.com/steve-care-software/products/diamonds/domain/genesis"
	"github.com/steve-care-software/products/libs/hash"
)

type genesis struct {
	hash      hash.Hash
	amount    hash.Hash
	seed      hash.Hash
	gen       domain_genesis.Genesis
	createdOn time.Time
}

func createGenesis(
	hash hash.Hash,
	amount hash.Hash,
	seed hash.Hash,
	gen domain_genesis.Genesis,
	createdOn time.Time,
) Genesis {
	out := genesis{
		hash:      hash,
		amount:    amount,
		seed:      seed,
		gen:       gen,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *genesis) Amount() hash.Hash {
	return obj.amount
}

// Seed returns the seed
func (obj *genesis) Seed() hash.Hash {
	return obj.seed
}

// Genesis returns the genesis
func (obj *genesis) Genesis() domain_genesis.Genesis {
	return obj.gen
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.createdOn
}
