package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/hash"
)

type accessible struct {
	immutable ImmutableAccessible
	mutable   MutableAccessible
}

func createAccessibleWithImmutable(
	immutable ImmutableAccessible,
) Accessible {
	return createAccessibleInternally(immutable, nil)
}

func createAccessibleWithMutable(
	mutable MutableAccessible,
) Accessible {
	return createAccessibleInternally(nil, mutable)
}

func createAccessibleInternally(
	immutable ImmutableAccessible,
	mutable MutableAccessible,
) Accessible {
	out := accessible{
		immutable: immutable,
		mutable:   mutable,
	}

	return &out
}

// ID returns the id
func (obj *accessible) ID() *uuid.UUID {
	if obj.IsImmutable() {
		return obj.Immutable().ID()
	}

	return obj.Mutable().ID()
}

// Hash returns the hash
func (obj *accessible) Hash() hash.Hash {
	if obj.IsImmutable() {
		return obj.Immutable().Hash()
	}

	return obj.Mutable().Hash()
}

// CreatedOn returns the cretion time
func (obj *accessible) CreatedOn() time.Time {
	if obj.IsImmutable() {
		return obj.Immutable().CreatedOn()
	}

	return obj.Mutable().CreatedOn()
}

// HasAccess true if there is an access, false otherwise
func (obj *accessible) HasAccess() bool {
	if obj.IsImmutable() {
		return obj.Immutable().HasAccess()
	}

	return obj.Mutable().HasAccess()
}

// HasAccess returns the access, if any
func (obj *accessible) Access() Access {
	if obj.IsImmutable() {
		return obj.Immutable().Access()
	}

	return obj.Mutable().Access()
}

// IsImmutable returns true if immutable, false otherwise
func (obj *accessible) IsImmutable() bool {
	return obj.immutable != nil
}

// Immutable returns the immutable, if any
func (obj *accessible) Immutable() ImmutableAccessible {
	return obj.immutable
}

// IsMutable returns true if mutable, false otherwise
func (obj *accessible) IsMutable() bool {
	return obj.mutable != nil
}

// Mutable returns the mutable, if any
func (obj *accessible) Mutable() MutableAccessible {
	return obj.mutable
}
