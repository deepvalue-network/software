package views

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends"
	"github.com/steve-care-software/products/libs/hash"
)

type genesis struct {
	hash   hash.Hash
	gen    spends.Genesis
	seed   string
	amount uint64
}

func createGenesis(
	hash hash.Hash,
	gen spends.Genesis,
	seed string,
	amount uint64,
) Genesis {
	out := genesis{
		hash:   hash,
		gen:    gen,
		seed:   seed,
		amount: amount,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Genesis returns the genesis
func (obj *genesis) Genesis() spends.Genesis {
	return obj.gen
}

// Seed returns the seed
func (obj *genesis) Seed() string {
	return obj.seed
}

// Amount returns the amount
func (obj *genesis) Amount() uint64 {
	return obj.amount
}
