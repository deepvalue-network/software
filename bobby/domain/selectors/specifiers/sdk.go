package specifiers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	identifiersBuilder := NewIdentifiersBuilder()
	return createBuilder(identifiersBuilder)
}

// NewIdentifiersBuilder creates a new identifiers builder
func NewIdentifiersBuilder() IdentifiersBuilder {
	hashAdapter := hash.NewAdapter()
	return createIdentifiersBuilder(hashAdapter)
}

// NewIdentifierBuilder creates a new identifier builder
func NewIdentifierBuilder() IdentifierBuilder {
	return createIdentifierBuilder()
}

// NewComparerBuilder creates a new comparer builder
func NewComparerBuilder() ComparerBuilder {
	hashAdapter := hash.NewAdapter()
	return createComparerBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// Builder represents a specifier builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier Identifier) Builder
	WithIdentifiers(identifiers []Identifier) Builder
	Now() (Specifier, error)
}

// Specifier represents a specifier
type Specifier interface {
	Hash() hash.Hash
	IsIdentifier() bool
	Identifier() Identifier
	IsIdentifiers() bool
	Identifiers() Identifiers
}

// IdentifiersBuilder represents an identifiers builder
type IdentifiersBuilder interface {
	Create() IdentifiersBuilder
	WithIdentifiers(identifiers []Identifier) IdentifiersBuilder
	Now() (Identifiers, error)
}

// Identifiers represents an identifiers
type Identifiers interface {
	Hash() hash.Hash
	All() []Identifier
}

// IdentifierBuilder represents an identifier builder
type IdentifierBuilder interface {
	Create() IdentifierBuilder
	WithElement(element Element) IdentifierBuilder
	WithComparer(comparer Comparer) IdentifierBuilder
	Now() (Identifier, error)
}

// Identifier represents a selector identifier
type Identifier interface {
	Hash() hash.Hash
	IsElement() bool
	Element() Element
	IsComparer() bool
	Comparer() Comparer
}

// ComparerBuilder represents a comparer builder
type ComparerBuilder interface {
	Create() ComparerBuilder
	WithFirst(first Identifier) ComparerBuilder
	WithSecond(second Identifier) ComparerBuilder
	IsAnd() ComparerBuilder
	Now() (Comparer, error)
}

// Comparer represents an identifier comparer
type Comparer interface {
	Hash() hash.Hash
	First() Identifier
	Second() Identifier
	IsAnd() bool
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithHash(hash hash.Hash) ElementBuilder
	WithID(id *uuid.UUID) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	IsHashPtr() bool
	HashPtr() *hash.Hash
	IsID() bool
	ID() *uuid.UUID
}
