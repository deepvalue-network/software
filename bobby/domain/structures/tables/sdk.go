package tables

import (
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
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
