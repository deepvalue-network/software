package governments

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

type government struct {
	hash    hash.Hash
	id      *uuid.UUID
	current Content
	holders shareholders.ShareHolders
	prev    Government
}

func createGovernment(
	hash hash.Hash,
	current Content,
	holders shareholders.ShareHolders,
	id *uuid.UUID,
) Government {
	return createGovernmentInternally(hash, current, holders, id, nil)
}

func createGovernmentWithPrevious(
	hash hash.Hash,
	current Content,
	holders shareholders.ShareHolders,
	prev Government,
) Government {
	return createGovernmentInternally(hash, current, holders, nil, prev)
}

func createGovernmentInternally(
	hash hash.Hash,
	current Content,
	holders shareholders.ShareHolders,
	id *uuid.UUID,
	prev Government,
) Government {
	out := government{
		hash:    hash,
		current: current,
		holders: holders,
		id:      id,
		prev:    prev,
	}

	return &out
}

// Hash returns the hash
func (obj *government) Hash() hash.Hash {
	return obj.hash
}

// ShareHolders returns the shareholders
func (obj *government) ShareHolders() shareholders.ShareHolders {
	return obj.holders
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
