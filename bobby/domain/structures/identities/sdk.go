package identities

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithResource(res resources.Mutable) Builder
	WithKey(key hash.Hash) Builder
	WithName(name string) Builder
	OnGraphbase(graphbase graphbases.Graphbase) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Resource() resources.Mutable
	Graphbase() graphbases.Graphbase
	Key() hash.Hash
	Name() string
}
