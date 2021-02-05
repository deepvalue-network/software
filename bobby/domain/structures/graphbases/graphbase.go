package graphbases

import (
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/resources"
)

type graphbase struct {
	resource resources.Accessible
	metaData string
	chain    chains.Chain
	parent   resources.Accessible
}

func createGraphbase(
	resource resources.Accessible,
	metaData string,
	chain chains.Chain,
) Graphbase {
	return createGraphbaseInternally(resource, metaData, chain, nil)
}

func createGraphbaseWithParent(
	resource resources.Accessible,
	metaData string,
	chain chains.Chain,
	parent resources.Accessible,
) Graphbase {
	return createGraphbaseInternally(resource, metaData, chain, parent)
}

func createGraphbaseInternally(
	resource resources.Accessible,
	metaData string,
	chain chains.Chain,
	parent resources.Accessible,
) Graphbase {
	out := graphbase{
		resource: resource,
		metaData: metaData,
		chain:    chain,
		parent:   parent,
	}

	return &out
}

// Resource returns the resource
func (obj *graphbase) Resource() resources.Accessible {
	return obj.resource
}

// MetaData returns the metaData table name
func (obj *graphbase) MetaData() string {
	return obj.metaData
}

// Chain returns the chain
func (obj *graphbase) Chain() chains.Chain {
	return obj.chain
}

// HasParent returns true if there is a parent, false otherwise
func (obj *graphbase) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *graphbase) Parent() resources.Accessible {
	return obj.parent
}
