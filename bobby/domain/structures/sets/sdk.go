package sets

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	resourceBuilder := resources.NewBuilder()
	return createBuilder(hashAdapter, resourceBuilder)
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementsBuilder(hashAdapter)
}

// Builder represents the set builder
type Builder interface {
	Create() Builder
	WithSchema(schema schemas.Schema) Builder
	WithElements(elements Elements) Builder
	WithName(name string) Builder
	OnGraphbase(graphbase graphbases.Graphbase) Builder
	Now() (Set, error)
}

// Set represents a set of elements
type Set interface {
	Resource() resources.Resource
	Graphbase() graphbases.Graphbase
	Schema() schemas.Schema
	Elements() Elements
	Name() string
}

// ElementsBuilder represents an elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithRanked(ranked map[uint]resources.Immutable) ElementsBuilder
	WithUnranked(unranked []resources.Immutable) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents a set of elements
type Elements interface {
	Hash() hash.Hash
	IsUnique() bool
	IsRanked() bool
	Ranked() RankedElements
	IsUnranked() bool
	UnRanked() UnrankedElements
}

// UnrankedElements represents unranked elements
type UnrankedElements interface {
	Hash() hash.Hash
	All() []resources.Immutable
	IsEmpty() bool
	IsUnique() bool
}

// RankedElements represents a ranked elements
type RankedElements interface {
	Hash() hash.Hash
	All() map[uint]resources.Immutable
	IsEmpty() bool
	IsUnique() bool
}
