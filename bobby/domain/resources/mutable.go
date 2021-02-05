package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

type mutable struct {
	immutable Immutable
	parent    Mutable
}

func createMutable(
	immutable Immutable,
) Mutable {
	return createMutableInternally(immutable, nil)
}

func createMutableWithParent(
	immutable Immutable,
	parent Mutable,
) Mutable {
	return createMutableInternally(immutable, parent)
}

func createMutableInternally(
	immutable Immutable,
	parent Mutable,
) Mutable {
	out := mutable{
		immutable: immutable,
		parent:    parent,
	}

	return &out
}

// ID returns the id
func (obj *mutable) ID() *uuid.UUID {
	return obj.immutable.ID()
}

// Hash returns the hash
func (obj *mutable) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// CreatedOn returns the cretion time
func (obj *mutable) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasParent returns true if there is a parent, false otherwise
func (obj *mutable) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *mutable) Parent() Mutable {
	return obj.parent
}
