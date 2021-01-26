package entities

import (
	"time"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

type mutable struct {
    hash hash.Hash
    createdOn time.Time
    lastUpdatedOn time.Time
}

func createMutable(hash hash.Hash, createdOn time.Time, lastUpdatedOn time.Time) Mutable {
    out := mutable{
        hash: hash,
        createdOn: createdOn,
        lastUpdatedOn: lastUpdatedOn,
    }

    return &out
}

// Hash returns the hash
func (obj *mutable) Hash() hash.Hash {
    return obj.hash
}

// CreatedOn returns the creation time
func (obj *mutable) CreatedOn() time.Time {
    return obj.createdOn
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *mutable) LastUpdatedOn() time.Time {
    return obj.lastUpdatedOn
}
