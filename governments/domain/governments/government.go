package governments

import (
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

type government struct {
	hash    hash.Hash
	id      *uuid.UUID
	current Content
	prev    Government
}

func createGovernment(
	hash hash.Hash,
	id *uuid.UUID,
	current Content,
) Government {
	return createGovernmentInternally(hash, current, id, nil)
}

func createGovernmentWithPrevious(
	hash hash.Hash,
	current Content,
	prev Government,
) Government {
	return createGovernmentInternally(hash, current, nil, prev)
}

func createGovernmentInternally(
	hash hash.Hash,
	current Content,
	id *uuid.UUID,
	prev Government,
) Government {
	out := government{
		hash:    hash,
		current: current,
		id:      id,
		prev:    prev,
	}

	return &out
}

// Hash returns the hash
func (obj *government) Hash() hash.Hash {
	return obj.hash
}

// ID returns the identifier
func (obj *government) ID() *uuid.UUID {
	if obj.prev != nil {
		return obj.prev.ID()
	}

	return obj.id
}

// Current returns the current government content
func (obj *government) Current() Content {
	return obj.current
}

// HasPrevious returns true if there is a previous government, false otherwise
func (obj *government) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous government, if any
func (obj *government) Previous() Government {
	return obj.prev
}
