package tables

import (
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	resourceBuilder := resources.NewBuilder()
	return createBuilder(hashAdapter, resourceBuilder)
}

// Builder represents the table builder
type Builder interface {
	Create() Builder
	WithSchema(schema schemas.Schema) Builder
	OnGraphbase(graphbase graphbases.Graphbase) Builder
	OnChain(chain chains.Chain) Builder
	Now() (Table, error)
}

// Table represents a table
type Table interface {
	Resource() resources.Resource
	Graphbase() graphbases.Graphbase
	Schema() schemas.Schema
	HasChain() bool
	Chain() chains.Chain
}
