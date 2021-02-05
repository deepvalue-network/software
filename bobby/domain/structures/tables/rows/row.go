package rows

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/elements"
)

type row struct {
	resource resources.Resource
	elements elements.Elements
	table    tables.Table
}

func createRow(
	resource resources.Resource,
	elements elements.Elements,
	table tables.Table,
) Row {
	out := row{
		resource: resource,
		elements: elements,
		table:    table,
	}

	return &out
}

// Resource returns the resource
func (obj *row) Resource() resources.Resource {
	return obj.resource
}

// Elements returns the elements
func (obj *row) Elements() elements.Elements {
	return obj.elements
}

// OnTable returns the table
func (obj *row) OnTable() tables.Table {
	return obj.table
}
