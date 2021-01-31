package structures

import (
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/identities"
)

type structure struct {
	graph    graphbases.Graphbase
	identity identities.Identity
	tb       Table
	set      Set
}

func createStructureWithGraph(
	graph graphbases.Graphbase,
) Structure {
	return createStructureInternally(graph, nil, nil, nil)
}

func createStructureWithIdentity(
	identity identities.Identity,
) Structure {
	return createStructureInternally(nil, identity, nil, nil)
}

func createStructureWithTable(
	tb Table,
) Structure {
	return createStructureInternally(nil, nil, tb, nil)
}

func createStructureWithSet(
	set Set,
) Structure {
	return createStructureInternally(nil, nil, nil, set)
}

func createStructureInternally(
	graph graphbases.Graphbase,
	identity identities.Identity,
	tb Table,
	set Set,
) Structure {
	out := structure{
		graph:    graph,
		identity: identity,
		tb:       tb,
		set:      set,
	}

	return &out
}

// IsGraphbase returns true if there is a graphbase, false otherwise
func (obj *structure) IsGraphbase() bool {
	return obj.graph != nil
}

// Graphbase returns the graphbase, if any
func (obj *structure) Graphbase() graphbases.Graphbase {
	return obj.graph
}

// IsIdentity returns true if there is an identity, false otherwise
func (obj *structure) IsIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *structure) Identity() identities.Identity {
	return obj.identity
}

// IsTable returns true if there is a table, false otherwise
func (obj *structure) IsTable() bool {
	return obj.tb != nil
}

// Table returns the table, if any
func (obj *structure) Table() Table {
	return obj.tb
}

// IsSet returns true if there is a set, false otherwise
func (obj *structure) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *structure) Set() Set {
	return obj.set
}
