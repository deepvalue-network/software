package identities

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/libs/hash"
)

type identity struct {
	resource  resources.Mutable
	graphbase graphbases.Graphbase
	key       hash.Hash
	name      string
}

func createIdentity(
	resource resources.Mutable,
	graphbase graphbases.Graphbase,
	key hash.Hash,
	name string,
) Identity {
	out := identity{
		resource:  resource,
		graphbase: graphbase,
		key:       key,
		name:      name,
	}

	return &out
}

// Resource returns the resource
func (obj *identity) Resource() resources.Mutable {
	return obj.resource
}

// Graphbase returns the graphbase
func (obj *identity) Graphbase() graphbases.Graphbase {
	return obj.graphbase
}

// Key returns the key
func (obj *identity) Key() hash.Hash {
	return obj.key
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}
