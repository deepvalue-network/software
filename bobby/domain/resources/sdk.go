package resources

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	immutableBuilder := NewImmutableBuilder()
	mutableBuilder := NewMutableBuilder()
	return createBuilder(immutableBuilder, mutableBuilder)
}

// NewMutableBuilder creates a new mutable builder instance
func NewMutableBuilder() MutableBuilder {
	immutableBuilder := NewImmutableBuilder()
	return createMutableBuilder(immutableBuilder)
}

// NewImmutableBuilder creates a new immutable builder instance
func NewImmutableBuilder() ImmutableBuilder {
	return createImmutableBuilder()
}

// NewAccessibleBuilder creates a new accessible builder instance
func NewAccessibleBuilder() AccessibleBuilder {
	return createAccessibleBuilder()
}

// NewMutableAccessibleBuilder creates a new mutable accessible builder instance
func NewMutableAccessibleBuilder() MutableAccessibleBuilder {
	mutableBuilder := NewMutableBuilder()
	return createMutableAccessibleBuilder(mutableBuilder)
}

// NewImmutableAccessibleBuilder creates a new immutable accessible builder instance
func NewImmutableAccessibleBuilder() ImmutableAccessibleBuilder {
	immutableBuilder := NewImmutableBuilder()
	return createImmutableAccessibleBuilder(immutableBuilder)
}

// NewAccessBuilder creates a new access builder instance
func NewAccessBuilder() AccessBuilder {
	return createAccessBuilder()
}

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithAccessible(accessible Accessible) Builder
	WithImmutable(immutable Immutable) Builder
	WithMutable(mutable Mutable) Builder
	Now() (Resource, error)
}

// Resource represents an immutable or mutable resource
type Resource interface {
	Immutable
	IsCompatible(accessible Accessible) bool
	IsImmutable() bool
	Immutable() Immutable
	IsMutable() bool
	Mutable() Mutable
}

// ImmutableBuilder represents an immutable builder
type ImmutableBuilder interface {
	Create() ImmutableBuilder
	WithHash(hash hash.Hash) ImmutableBuilder
	Now() (Immutable, error)
}

// Immutable represents an immutable resource
type Immutable interface {
	ID() *uuid.UUID
	Hash() hash.Hash
	CreatedOn() time.Time
}

// MutableBuilder represents a mutable builder
type MutableBuilder interface {
	Create() MutableBuilder
	WithHash(hash hash.Hash) MutableBuilder
	WithParent(parent Mutable) MutableBuilder
	Now() (Mutable, error)
}

// Mutable represents a mutable resource
type Mutable interface {
	Immutable
	HasParent() bool
	Parent() Mutable
}

// AccessibleBuilder represents an accessible builder
type AccessibleBuilder interface {
	Create() AccessibleBuilder
	WithImmutable(immutable ImmutableAccessible) AccessibleBuilder
	WithMutable(mutable MutableAccessible) AccessibleBuilder
	Now() (Accessible, error)
}

// Accessible represents an accessible resource
type Accessible interface {
	ImmutableAccessible
	IsImmutable() bool
	Immutable() ImmutableAccessible
	IsMutable() bool
	Mutable() MutableAccessible
}

// ImmutableAccessibleBuilder represents an immutable accessible builder
type ImmutableAccessibleBuilder interface {
	Create() ImmutableAccessibleBuilder
	WithHash(hash hash.Hash) ImmutableAccessibleBuilder
	WithAccess(access Access) ImmutableAccessibleBuilder
	Now() (ImmutableAccessible, error)
}

// ImmutableAccessible represents an immutable accessible resource
type ImmutableAccessible interface {
	Immutable
	HasAccess() bool
	Access() Access
}

// MutableAccessibleBuilder represents a mutable accessible builder
type MutableAccessibleBuilder interface {
	Create() MutableAccessibleBuilder
	WithHash(hash hash.Hash) MutableAccessibleBuilder
	WithParent(parent Mutable) MutableAccessibleBuilder
	WithAccess(access Access) MutableAccessibleBuilder
	Now() (MutableAccessible, error)
}

// MutableAccessible represents a mutable accessible resource
type MutableAccessible interface {
	Mutable
	HasAccess() bool
	Access() Access
}

// AccessBuilder represents an access builder
type AccessBuilder interface {
	Create() AccessBuilder
	WithResource(res Mutable) AccessBuilder
	WithOwners(owners []*uuid.UUID) AccessBuilder
	WithEncryptionPubKey(pubKey public.Key) AccessBuilder
	Now() (Access, error)
}

// Access represents a list of pubkeys that can access data
type Access interface {
	Resource() Mutable
	Owners() []*uuid.UUID
	IsEncrypted() bool
	Encrypted() public.Key
}
