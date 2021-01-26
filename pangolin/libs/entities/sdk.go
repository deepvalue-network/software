package entities

import (
	"time"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

// NewImmutableBuilder creates a new immutable builder
func NewImmutableBuilder() ImmutableBuilder {
	return createImmutableBuilder()
}

// NewMutableBuilder creates a new mutable builder
func NewMutableBuilder() MutableBuilder {
	return createMutableBuilder()
}

// ImmutableBuilder represents an immutable builder
type ImmutableBuilder interface {
	Create() ImmutableBuilder
	WithHash(hash hash.Hash) ImmutableBuilder
	CreatedOn(createdOn *time.Time) ImmutableBuilder
	Now() (Immutable, error)
}

// Immutable represents an immutable entity
type Immutable interface {
	Hash() hash.Hash
	CreatedOn() time.Time
}

// MutableBuilder represents a mutable builder
type MutableBuilder interface {
	Create() MutableBuilder
	WithHash(hash hash.Hash) MutableBuilder
	CreatedOn(createdOn *time.Time) MutableBuilder
	LastUpdatedOn(lastUpdatedOn *time.Time) MutableBuilder
	Now() (Mutable, error)
}

// Mutable represents a mutable entity
type Mutable interface {
	Immutable
	LastUpdatedOn() time.Time
}
