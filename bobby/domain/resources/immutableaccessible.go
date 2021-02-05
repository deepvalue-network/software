package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

type immutableAccessible struct {
	immutable Immutable
	access    Access
}

func createImmutableAccessible(
	immutable Immutable,
) ImmutableAccessible {
	return createImmutableAccessibleInternally(immutable, nil)
}

func createImmutableAccessibleWithAccess(
	immutable Immutable,
	access Access,
) ImmutableAccessible {
	return createImmutableAccessibleInternally(immutable, access)
}

func createImmutableAccessibleInternally(
	immutable Immutable,
	access Access,
) ImmutableAccessible {
	out := immutableAccessible{
		immutable: immutable,
		access:    access,
	}

	return &out
}

// ID returns the id
func (obj *immutableAccessible) ID() *uuid.UUID {
	return obj.immutable.ID()
}

// Hash returns the hash
func (obj *immutableAccessible) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// CreatedOn returns the cretion time
func (obj *immutableAccessible) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasAccess returns true if there is an access, false otherwise
func (obj *immutableAccessible) HasAccess() bool {
	return obj.access != nil
}

// Access returns the access, if any
func (obj *immutableAccessible) Access() Access {
	return obj.access
}
