package transfers

import (
	"time"

	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	origin    hash.Hash
	amount    hash.Hash
	owner     []hash.Hash
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	origin hash.Hash,
	amount hash.Hash,
	owner []hash.Hash,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		origin:    origin,
		amount:    amount,
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
func (obj *content) Origin() hash.Hash {
	return obj.origin
}

// Amount returns the amount
func (obj *content) Amount() hash.Hash {
	return obj.amount
}

// Owner returns the owner
func (obj *content) Owner() []hash.Hash {
	return obj.owner
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
