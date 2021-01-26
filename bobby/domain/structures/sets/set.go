package sets

import (
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
)

type set struct {
	resource  resources.Resource
	graphbase graphbases.Graphbase
	schema    schemas.Schema
	elements  Elements
	name      string
}

func createSet(
	resource resources.Resource,
	graphbase graphbases.Graphbase,
	schema schemas.Schema,
	elements Elements,
	name string,
) Set {
	out := set{
		resource:  resource,
		graphbase: graphbase,
		schema:    schema,
		elements:  elements,
		name:      name,
	}

	return &out
}

// Resource returns the resource
func (obj *set) Resource() resources.Resource {
	return obj.resource
}

// Graphbase returns the graphbase
func (obj *set) Graphbase() graphbases.Graphbase {
	return obj.graphbase
}

// Schema returns the schema
func (obj *set) Schema() schemas.Schema {
	return obj.schema
}

// Elements returns the elements
func (obj *set) Elements() Elements {
	return obj.elements
}

// Name returns the name
func (obj *set) Name() string {
	return obj.name
}
