package schemas

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables"
)

type schema struct {
	resource         resources.Accessible
	name             string
	table            tables.Table
	isUniqueElements bool
}

func createSchema(
	resource resources.Accessible,
	name string,
	table tables.Table,
	isUniqueElements bool,
) Schema {
	out := schema{
		resource:         resource,
		name:             name,
		table:            table,
		isUniqueElements: isUniqueElements,
	}

	return &out
}

// Resource returns the resource
func (obj *schema) Resource() resources.Accessible {
	return obj.resource
}

// Name returns the name
func (obj *schema) Name() string {
	return obj.name
}

// Table returns the table
func (obj *schema) Table() tables.Table {
	return obj.table
}

// IsUniqueElements returns true if the elements are unique, false otherwise
func (obj *schema) IsUniqueElements() bool {
	return obj.isUniqueElements
}
