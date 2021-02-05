package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

type immutable struct {
	id        *uuid.UUID
	hash      hash.Hash
	createdOn time.Time
}

func createImmutable(
	id *uuid.UUID,
	hash hash.Hash,
	createdOn time.Time,
) Immutable {
	out := immutable{
		id:        id,
		hash:      hash,
		createdOn: createdOn,
	}

	return &out
}

// ID returns the id
func (obj *immutable) ID() *uuid.UUID {
	return obj.id
}

// Hash returns the hash
func (obj *immutable) Hash() hash.Hash {
	return obj.hash
}

// CreatedOn returns the cretion time
func (obj *immutable) CreatedOn() time.Time {
	return obj.createdOn
}
