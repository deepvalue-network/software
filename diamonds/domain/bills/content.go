package bills

import (
	"time"

	"github.com/steve-care-software/products/libs/hash"
)

type content struct {
	hash      hash.Hash
	origin    Origin
	amount    hash.Hash
	seed      hash.Hash
	owner     []hash.Hash
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	origin Origin,
	amount hash.Hash,
	seed hash.Hash,
	owner []hash.Hash,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		origin:    origin,
		amount:    amount,
		seed:      seed,
		owner:     owner,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Origin returns the origin
func (obj *content) Origin() Origin {
	return obj.origin
}

// Amount returns the amount
func (obj *content) Amount() hash.Hash {
	return obj.amount
}

// Seed returns the seed
func (obj *content) Seed() hash.Hash {
	return obj.seed
}

// Owner returns the owner
func (obj *content) Owner() []hash.Hash {
	return obj.owner
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
