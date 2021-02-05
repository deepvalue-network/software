package tables

import (
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
)

type table struct {
	resource  resources.Resource
	graphbase graphbases.Graphbase
	schema    schemas.Schema
	chain     chains.Chain
}

func createTable(
	resource resources.Resource,
	graphbase graphbases.Graphbase,
	schema schemas.Schema,
) Table {
	return createTableInternally(resource, graphbase, schema, nil)
}

func createTableWithChain(
	resource resources.Resource,
	graphbase graphbases.Graphbase,
	schema schemas.Schema,
	chain chains.Chain,
) Table {
	return createTableInternally(resource, graphbase, schema, chain)
}

func createTableInternally(
	resource resources.Resource,
	graphbase graphbases.Graphbase,
	schema schemas.Schema,
	chain chains.Chain,
) Table {
	out := table{
		resource:  resource,
		graphbase: graphbase,
		schema:    schema,
		chain:     chain,
	}

	return &out
}

// Resource returns the resource
func (obj *table) Resource() resources.Resource {
	return obj.resource
}

// Graphbase returns the graphbase
func (obj *table) Graphbase() graphbases.Graphbase {
	return obj.graphbase
}

// Schema returns the schema
func (obj *table) Schema() schemas.Schema {
	return obj.schema
}

// HasChain returns true if there is a chain, false otherwise
func (obj *table) HasChain() bool {
	return obj.chain != nil
}

// Chain returns the chain, if any
func (obj *table) Chain() chains.Chain {
	return obj.chain
}
