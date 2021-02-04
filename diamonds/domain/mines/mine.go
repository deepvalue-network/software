package mines

import (
	"time"

	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hashtree"
)

type mine struct {
	hash      hash.Hash
	chain     chains.Chain
	diamonds  hashtree.HashTree
	createdOn time.Time
}

func createMine(
	hash hash.Hash,
	chain chains.Chain,
	diamonds hashtree.HashTree,
	createdOn time.Time,
) Mine {
	out := mine{
		hash:      hash,
		chain:     chain,
		diamonds:  diamonds,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *mine) Hash() hash.Hash {
	return obj.hash
}

// Chain returns the chain
func (obj *mine) Chain() chains.Chain {
	return obj.chain
}

// Diamonds returns the diamonds
func (obj *mine) Diamonds() hashtree.HashTree {
	return obj.diamonds
}

// CreatedOn returns the creation time
func (obj *mine) CreatedOn() time.Time {
	return obj.createdOn
}
