package access

import (
	access "github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an access transaction builder
type Builder interface {
	Create() Builder
	WithResources(resources selectors.Selector) Builder
	Add(add access.Access) Builder
	Remove() Builder
	Now() (Transaction, error)
}

// Transaction represents an access transaction
type Transaction interface {
	Hash() hash.Hash
	Resources() selectors.Selector
	Content() Content
}

// Content represents an entity transaction content
type Content interface {
	IsRemove() bool
	IsAdd() bool
	Add() access.Access
}
