package elements

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/values"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	resourceBuilder := resources.NewBuilder()
	return createElementBuilder(hashAdapter, resourceBuilder)
}

// Builder represents the elemnts builder
type Builder interface {
	Create() Builder
	WithElements(elements []Element) Builder
	Now() (Elements, error)
}

// Elements represents the elements
type Elements interface {
	Hash() hash.Hash
	All() []Element
	IsEmpty() bool
	Fits(properties schemas.Properties) error
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithProperty(property schemas.Property) ElementBuilder
	WithValue(value values.Value) ElementBuilder
	Now() (Element, error)
}

// Element represents an element in a row
type Element interface {
	Resource() resources.Resource
	Property() schemas.Property
	Value() values.Value
}
