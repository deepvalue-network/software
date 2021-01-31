package structures

import (
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/identities"
)

type content struct {
	graph    graphbases.Graphbase
	identity identities.Identity
	tb       Table
	set      Set
}

func createContentWithGraph(
	graph graphbases.Graphbase,
) Content {
	return createContentInternally(graph, nil, nil, nil)
}

func createContentWithIdentity(
	identity identities.Identity,
) Content {
	return createContentInternally(nil, identity, nil, nil)
}

func createContentWithTable(
	tb Table,
) Content {
	return createContentInternally(nil, nil, tb, nil)
}

func createContentWithSet(
	set Set,
) Content {
	return createContentInternally(nil, nil, nil, set)
}

func createContentInternally(
	graph graphbases.Graphbase,
	identity identities.Identity,
	tb Table,
	set Set,
) Content {
	out := content{
		graph:    graph,
		identity: identity,
		tb:       tb,
		set:      set,
	}

	return &out
}

// IsGraphbase returns true if there is a graphbase, false otherwise
func (obj *content) IsGraphbase() bool {
	return obj.graph != nil
}

// Graphbase returns the graphbase, if any
func (obj *content) Graphbase() graphbases.Graphbase {
	return obj.graph
}

// IsIdentity returns true if there is an identity, false otherwise
func (obj *content) IsIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *content) Identity() identities.Identity {
	return obj.identity
}

// IsTable returns true if there is a table, false otherwise
func (obj *content) IsTable() bool {
	return obj.tb != nil
}

// Table returns the table, if any
func (obj *content) Table() Table {
	return obj.tb
}

// IsSet returns true if there is a set, false otherwise
func (obj *content) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *content) Set() Set {
	return obj.set
}
