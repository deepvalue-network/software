package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

type mutableAccessible struct {
	mutable Mutable
	access  Access
}

func createMutableAccessible(
	mutable Mutable,
) MutableAccessible {
	return createMutableAccessibleInternally(mutable, nil)
}

func createMutableAccessibleWithAccess(
	mutable Mutable,
	access Access,
) MutableAccessible {
	return createMutableAccessibleInternally(mutable, access)
}

func createMutableAccessibleInternally(
	mutable Mutable,
	access Access,
) MutableAccessible {
	out := mutableAccessible{
		mutable: mutable,
		access:  access,
	}

	return &out
}

// ID returns the id
func (obj *mutableAccessible) ID() *uuid.UUID {
	return obj.mutable.ID()
}

// Hash returns the hash
func (obj *mutableAccessible) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// CreatedOn returns the cretion time
func (obj *mutableAccessible) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// HasParent returns true if there is a parent, false otherwise
func (obj *mutableAccessible) HasParent() bool {
	return obj.mutable.HasParent()
}

// Parent returns the parent, if any
func (obj *mutableAccessible) Parent() Mutable {
	return obj.mutable.Parent()
}

// HasAccess returns true if there is an access, false otherwise
func (obj *mutableAccessible) HasAccess() bool {
	return obj.access != nil
}

// Access returns the access, if any
func (obj *mutableAccessible) Access() Access {
	return obj.access
}
