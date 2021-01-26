package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/hash"
)

type resource struct {
	mutable   Mutable
	immutable Immutable
}

func createResourceWithMutable(
	mutable Mutable,
) Resource {
	return createResourceInternally(mutable, nil)
}

func createResourceWithImmutable(
	immutable Immutable,
) Resource {
	return createResourceInternally(nil, immutable)
}

func createResourceInternally(
	mutable Mutable,
	immutable Immutable,
) Resource {
	out := resource{
		mutable:   mutable,
		immutable: immutable,
	}

	return &out
}

// ID returns the id
func (obj *resource) ID() *uuid.UUID {
	if obj.IsImmutable() {
		return obj.Immutable().ID()
	}

	return obj.Mutable().ID()
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	if obj.IsImmutable() {
		return obj.Immutable().Hash()
	}

	return obj.Mutable().Hash()
}

// CreatedOn returns the cretion time
func (obj *resource) CreatedOn() time.Time {
	if obj.IsImmutable() {
		return obj.Immutable().CreatedOn()
	}

	return obj.Mutable().CreatedOn()
}

// IsCompatible returns true if compatible, false otherwise
func (obj *resource) IsCompatible(accessible Accessible) bool {
	if obj.IsImmutable() {
		return accessible.IsImmutable()
	}

	return accessible.IsMutable()
}

// IsImmutable returns true if the resource is immutable, false otherwise
func (obj *resource) IsImmutable() bool {
	return obj.immutable != nil
}

// Immutable returns the immutable, if any
func (obj *resource) Immutable() Immutable {
	return obj.immutable
}

// IsMutable returns true if the resource is mutable, false otherwise
func (obj *resource) IsMutable() bool {
	return obj.mutable != nil
}

// Mutable returns the mutable, if any
func (obj *resource) Mutable() Mutable {
	return obj.mutable
}
