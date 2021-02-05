package specifiers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

type element struct {
	hash    hash.Hash
	hashPtr *hash.Hash
	id      *uuid.UUID
}

func createElementWithHash(
	hash hash.Hash,
	hashPtr *hash.Hash,
) Element {
	return createElementInternally(hash, hashPtr, nil)
}

func createElementWithID(
	hash hash.Hash,
	id *uuid.UUID,
) Element {
	return createElementInternally(hash, nil, id)
}

func createElementInternally(
	hash hash.Hash,
	hashPtr *hash.Hash,
	id *uuid.UUID,
) Element {
	out := element{
		hash:    hash,
		hashPtr: hashPtr,
		id:      id,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// IsHashPtr returns true if there is an hash pointer, false otherwise
func (obj *element) IsHashPtr() bool {
	return obj.hashPtr != nil
}

// HashPtr returns the hash pointer, if set
func (obj *element) HashPtr() *hash.Hash {
	return obj.hashPtr
}

// IsID returns true if there is an ID, false otherwise
func (obj *element) IsID() bool {
	return obj.id != nil
}

// ID returns the id, if set
func (obj *element) ID() *uuid.UUID {
	return obj.id
}
